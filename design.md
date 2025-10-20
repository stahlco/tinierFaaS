# IDEA

- FaaS-Plattform
- Only HTTP
- Only Upload/Deletion of a Function
- And Calling
- Only shell functions

#### How to implement?

- Reverse Proxy -> map[function-names]ip
- Management-Server: (Listens on :8080)
  - UploadFunc -> Calls ":8080/upload" creates new functionHandler (works like...)
  - DeleteFunc ->

### How to call/invoke a function?
- via reverse proxy?
- via management-service?

### Docker-Backend
- Just for shell -> cmd.Exec(./fn.sh) ...
- 
