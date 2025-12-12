## 서버 구성

client : 슝의 클라이언트입니다. (Nextjs) 
authwz : 슝의 인증 클라이언트입니다. (Nextjs) 
stream : 슝의 채팅서버입니다. (Elixir/Phoenix) 
ingest : 슝의 RTMP를 입력받는 서버입니다. (Go) 
transcode : RTMP를 후킹하여 Hls로 트랜스코딩하는 서버입니다. (Go) 
authn : 슝의 인증을 관리하는 서버입니다. (Nestjs) 
authz : 슝의 인가를 관리하는 Zanzibar입니다. (Rust) 

각 서버간 통신은 gRPC를 통해 이루어집니다.

pg : 슝의 메인 db입니다. (postgres) 
redis-edge : 슝의 메인 캐싱 서버입니다. (redis) 
redis-shared : 슝의 msa 메모리 공유를 위한 db입니다. (redis) 
redis-authz : 슝의 authz를 위한 db입니다. (redis) 
dynamo-chat : 슝의 채팅서버입니다. (redis) 