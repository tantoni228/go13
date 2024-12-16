# go13
go13
v арі
 chats-service.yml
 messages-service.yml
 schemas.yml
 swagger ymi
 users-service.yml
v chats-service
 v cmd
  v main
   main.go
  v messages-service-mock
   Dockerfile
   main.go
 v configs  
  example.yml 
 v internal
  v config
   config.go 
  v models
   chat.go
   errors.go
   member.go
   message.go
   role.go
  v rеро
   v messages
    messages.go 
   v postgres
    chats.go
    members.go
    roles.go 
  v service
   access.go
   chats.go
   interfaces.go
   roles.go
  v transport
   v rest
    v auth
     auth.go
    v handlers
     chats.go
     roles.go
    v mapper
     roles.go
    server.go
 v migrations
   000001 create chats table.down.sal
   000001_create_chats_table.up.sql
   000002 create roles table.down.sal
   000002_create_roles_table.up.sql
   000003 create members table.down.sql
   000003 create members table.up.sal
   000004 add is system to roles.down.sql
   000004_add_is_system_to_roles.up.sq|
   000005 create banned members table.down.sql
   000005 create banned members table.up.sql
   000006 rename can ban_users to can_ manage members.down.sql
   000006_rename_can_ban_users_to_can_manage_members.up.sql
 v nginx
  nginx.conf
 docker-compose.yml
v cmd
 v main
  main.go
v docker
 Dockerfile.chats-service
v message-service
 v cmd
  v main
   main.go
 v configs
  example.yml
  kafka-variables.env
 v internal
  v config
   config.go
  v models
   message.go
   errors.go
  v repo
   v postgres
    message.go
  v service
   message.go  
  v transport
   v rest
    v auth
     auth.go
    v handlers
     message.go
    v middlewares
     apply.go
     logging.go
    server.go
 v migrations
  000001_create_messages_table.down.sql
  000001_create_messages_table.up.sql
 docker-compose.yml
v pkg
 v logger
  logger.go
 v middlewares
  apply.go
  logging.go
 v ogen
  v chats-service
   oas_cfg_gen.go
   oas_client_gen.go
   oas_handlers_gen.go
   oas_interfaces_gen.go
   oas_json_gen.go
   oas_middleware_gen.go
   oas_operations_gen.go
   oas_parameters_gen.go
   oas_request_decoders_gen.go
   oas_request_encoders_gen.go
   oas_response_decoders_gen.go
   aas_response_encoders_gen.go
   oas_router_gen.go
   oas_schemas_gen.go
   oas_security_gen.go
   oas_server_gen.go
   oas validators_gen.go
  v messages-service
   oas_cfg_gen.go
   oas_client_gen.go
   oas_handlers_gen.go
   oas_interfaces_gen.go
   oas_json_gen.go
   oas_middleware_gen.go
   oas_operations_gen.go
   oas_parameters_gen.go
   oas_request_decoders_gen.go
   oas_request_encoders_gen.go
   oas_response_decoders_gen.go
   aas_response_encoders_gen.go
   oas_router_gen.go
   oas_schemas_gen.go
   oas_security_gen.go
   oas_server_gen.go
   oas validators_gen.go
  v users-service
   oas_cfg_gen.go
   oas_client_gen.go
   oas_handlers_gen.go
   oas_interfaces_gen.go
   oas_json_gen.go
   oas_middleware_gen.go
   oas_operations_gen.go
   oas_parameters_gen.go
   oas_request_decoders_gen.go
   oas_request_encoders_gen.go
   oas_response_decoders_gen.go
   aas_response_encoders_gen.go
   oas_router_gen.go
   oas_schemas_gen.go
   oas_security_gen.go
   oas_server_gen.go
   oas validators_gen.go
 v postgres
  config.go
  postgres.go
v user-service
 v cmd
  v main
   main.go
 v configs
 v internal
 v migrations