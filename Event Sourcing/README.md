# Event Sourcing

Instead of storing the actual data in a database, events are stored as a sequence, and the current state is rebuilt by replaying these events in order. Instead of updating DB records directly, each change is recorded as an immutable event.

### Event Storage Example

**Bank DB**: Instead of storing the current balance and changing them on deposit/withdrawal, we can store events. We can replay these events to get the Account Balance.

- `AccountOpened { accountID: 123, initialBalance: 0 }`
- `MoneyDeposited { accountID: 123, amount: 100 }`
- `MoneyWithdrawn { accountID: 123, amount: 50 }`

**New York Times**: They have stored each article and its changes as events. So, if anyone wants to look at an older version, they can show that. Otherwise, just replay all these events to recreate the article.

- `ArticleCreated`
- `HeadlineUpdated`
- `ImageAdded`
- `BylineUpdated`

### CDC (Change Data Capture)

CDC tools (like Debezium) watch your database and generate events every time something changes.

Example:

1. Let’s say you have made an update to the user table:
    - `UPDATE users SET email = "newexample@gmail.com" WHERE id = 1;`
  
2. CDC captures this as an event:
    ```json
    {
        "type": "UserEmailUpdated",
        "userId": 1,
        "newEmail": "new@example.com",
        "timestamp": "2025-04-07T08:00:00Z"
    }
    ```

3. This event can be consumed by other services (like notifying users or logging the changes) without directly touching the database again.

### Outbox Pattern

Instead of directly sending a message to another system as part of a database operation (e.g., during a CREATE, UPDATE, or DELETE action), the Outbox Pattern stores the message in the same transactional context as the database operation. A separate process or service then reads the messages from the outbox and sends them to the message broker.

#### Example Workflow:

1. **Write to outbox in a transaction**:
    - When a new order is created, the system creates two records in one transaction:
      - One in the orders table for the new order.
      - One in the outbox table for the event `OrderCreated` (including details about the order).
      - Ensures both the order and outbox records are saved together, so if the system crashes, both transactions are cancelled.

2. **Poll the outbox and send events**:
    - A background worker or polling service continuously checks the outbox table for unsent events.
    - The service reads an unsent event, sends it to the message broker, and then marks it as "sent" by updating the `sent_at` field.

3. **Message broker delivery**:
    - The event is sent to the message broker (like Kafka, RabbitMQ), where other services (e.g., inventory, shipping) can consume the event and act accordingly.

### Microservices

Actions recorded as events:

- `ItemAddedToCart { userId: 101, itemId: 432, quantity: 2 }`
- `ItemRemovedFromCart { userId: 101, itemId: 432 }`
- `CartCheckout { userId: 101, cartId: 77 }`

This system can also store a timeline of what happened, helping other services with:

- **InventoryService**: Knows which item to reduce and how much to reduce.
- **BillingService**: Knows what to charge.

Event sourcing could also be used for more examples, such as:

1. Ticket Booking System
2. Invoice Service
3. Learning Management System (CourseEnrolled, LessonLearned, QuizSubmitted, …)
4. Logistics and Fleet Management
5. Game Backend

### When to Use Event Sourcing

Event sourcing isn’t always the right default, even if it seems cool and powerful. Here’s a breakdown of when to use event sourcing:

- **You need full history and audibility**:
  - Healthcare (patient records, prescriptions)
  - Finance (transactions, payments)
  - Legal/Insurance (claims, policy changes)

- **The domain is highly collaborative or stateful**:
  - Google Docs, Overleaf, game servers, Figma
  - You need to track every change here.

- **You need to rebuild the state on demand**:
  - If the system’s state evolves or changes often, and you want to be able to replay events with new logic.
  - Change how balance or tax is calculated — just reprocess the event with new rules.

- **You have Event-Driven Microservices**:
  - Events can be used to notify other services (CQRS pattern).
  - Good for decoupling and async processing, and scaling write vs read paths.

### When Not to Use Event Sourcing

- **Simplicity is better**:
  - If all you need is a basic CRUD app, event sourcing is overkill.
  - A portal that just displays some data may not need this complexity.

- **No need for history**:
  - If you don’t need an audit or traceback.

- **Real-time performance is critical**:
  - Replaying millions of events to get the state may slow down reads unless we build snapshots — this adds complexity.

### Ask Yourself

| **Criteria**      | **Ask Yourself**                                |
|-------------------|------------------------------------------------|
| **Audibility**     | Do I need to know how this state came to be?   |
| **Debugging**      | Will I want to rewind or re-simulate the system? |
| **Regulatory**     | Will legal/compliance require a full trace?    |
| **Scalability**    | Will decoupled microservices benefit from events? |
| **Collaboration**  | Do users work on the same data and need history/undo? |
| **Change Frequency** | Will my business logic change often enough to play things differently? |
