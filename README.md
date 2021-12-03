# Better ux
- alias neo='./neo'

# Goals
    - Embrace top community packages as much as possible.
    - Unified fluent API.
    - No reflection, struct tags as much as possible.
    - Give users choice of whether to lean towards a more dynamic or a strict approach.
    - Should be suitable for both HTTP/GRPC APIs and long running workers.

# Components
    - Dependency Injection Engine
    - HTTP Engine [x]
    - GRPC Engine
    - Configuration Management Engine [x]
    - I18n Engine
    - Background Job/Worker engine (asynq probably)
    - Event/PubSub engine 
    - Storage Engine (ent) [x]
