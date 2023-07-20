# Service Saksi

## How to

### Checkout

```bash
git clone https://repo.mncinnovation.id/perindo/saksi/service-saksi
make init
```

### Run from Source

```bash
make run
```

### Build

```bash
make build
```

### Run

To run Service Saksi with default configuration use

```bash
./service-saksi
```

To run Service Saksi with configuration use

```bash
./service-saksi -c config.yaml
```

To run Service Saksi with environment variable

```bash
SERVER_PORT=8080 ./service-saksi
```

### Config and Environment Variable

You can run use YAML config file or environment variable. Here are the parameters.

| Config File | Environment Variable | Type | Default Value | Description |
|-------------|----------------------|------|---------------|-------------|
| server.port | SERVER_PORT | String | 8555 | Local machine TCP Port to bind the HTTP Server to |
| server.prefork | SERVER_PREFORK | Boolean | false | Prefork will spawn multiple Go processes listening on the same port |
| server.strict_routing | SERVER_STRICT_ROUTING | Boolean | false | When enabled, the router treats /foo and /foo/ as different |
| server.case_sensitive | SERVER_CASE_SENSITIVE | Boolean | false | When enabled, /Foo and /foo are different routes |
| server.body_limit | SERVER_BODY_LIMIT | Integer | 4194304 | Sets the maximum allowed size for a request body |
| server.concurrency | SERVER_CONCURRENCY | Integer | 262144 | Concurrency maximum number of concurrent connections |
| server.timeout.read | SERVER_TIMEOUT_READ | Integer | 5 | The amount of time to wait until an HTTP server read operation is cancelled |
| server.timeout.write | SERVER_TIMEOUT_WRITE | Integer | 10 | The amount of time to wait until an HTTP server write operation is cancelled |
| server.timeout.idle | SERVER_TIMEOUT_IDLE | Integer | 120 | The amount of time to wait until an IDLE HTTP session is closed |
| server.log_level | SERVER_LOG_LEVEL | String | debug | Log level, available value: `error`, `warning`, `info`, `debug` |
| redis.host | REDIS_HOST | String | localhost | The Redis IP Address to connect to |
| redis.port | REDIS_PORT | String | 6379 | The Redis Port to connect to |
| redis.max_connection | REDIS_MAX_CONNECTION | Integer | 80 | Redis maximum connection |
| redis.username | REDIS_USERNAME | String | | Redis username |
| redis.password | REDIS_PASSWORD | String | | Redis password |
| redis.database | REDIS_DATABASE | Integer | 0 | Redis database number |
| middleware.allows_origin | MIDDLEWARE_ALLOWS_ORIGIN | String | * | List of origins that allow for CORS |

### Flowchart

<!-- ```plantuml
@startuml
split
-[hidden]->
start
split again
-[hidden]->
(A)
end split
:Landing page app Saksi;
if (Already\nregistered\nSaksi?) then (yes)
  :Login app Saksi;
  if (Already\nregistered\nonline member?) then (yes)
    if (Wrong\npassword?) then (yes)
      :Reset password;
      repeat :Request OTP
      repeat while (Wrong\nOTP?) is (No) not (Yes)
      :Reset password success;
      stop
    else (no)
      end
    endif
  else (no)
    :Unregistered online member;
    if (Register online member?) then (Yes)
      :Registered online member;
      stop
    else (No)
      (A)
      detach;
    endif
  endif
else (no)
  end
endif
@enduml
``` -->

### Sequence Diagram

<!-- 1. Login Saksi success

```seqdiag
seqdiag {
  APP; service-saksi-management; be-service-member; be-service-auth;
  APP => be-service-auth [label = "phone_number, password", note = "GET\n/login", return = token]
  APP -> service-saksi-management [note = "GET\n/auth"]
  service-saksi-management => be-service-auth [note = "gRPC\nGetSessionServiceAuth", color = green, label = AuthorizationAuthServiceRequest, return = AuthorizationAuthServiceResponse]
  service-saksi-management => be-service-member [note = "gRPC\nGetDetailMemberByID", color = green, label = GetDetailMemberByIDRequest, return = GetDetailMemberResponse]
  APP <-- service-saksi-management [label = "saksi data"]
}
```

2. Register Saksi, already registered online member

```seqdiag
seqdiag {
  APP; service-saksi-management; be-service-member; be-service-auth;
  APP => be-service-auth [label = "phone_number, password", note = "GET\n/login", return = token]
  APP -> service-saksi-management [label = "selfie_image, ktp_image", note = "POST\n/saksi/validation"]
  service-saksi-management -> be-service-member [note = "gRPC\nGetDetailMemberByNo", color = green, label = GetDetailMemberByNoRequest]
  service-saksi-management <-- be-service-member [leftnote = "Save to DB table saksi_candidate", color = green, label = GetDetailMemberResponse]
  APP <-- service-saksi-management [leftnote = "Send QR Code to DPRT", label = "nik, full_name, phone_number, address, dprt_coordinator, qr_code"]
}
```

3. Register Saksi, unregistered online member

```seqdiag
seqdiag {
  SMS; APP; service-saksi-management; be-service-member; be-service-auth;
  APP -> be-service-auth [note = "GET\n/otp", label = "number"]
  SMS -> APP [label = "OTP"]
  APP => be-service-auth [note = "POST\n/verifyotp", label = "number, otp, password", return = token]
  APP => be-service-member [note = "POST\n/register"]
  APP -> service-saksi-management [label = "selfie_image, ktp_image", note = "POST\n/saksi/validation"]
  service-saksi-management -> be-service-member [note = "gRPC\nGetDetailMemberByNo", color = green, label = GetDetailMemberByNoRequest]
  service-saksi-management <-- be-service-member [leftnote = "Save to DB table saksi_candidate", color = green, label = GetDetailMemberResponse]
  APP <-- service-saksi-management [leftnote = "Send QR Code to DPRT", label = "nik, full_name, phone_number, address, dprt_coordinator, qr_code"]
}
```

4. Check request register Saksi from DPRT side

```seqdiag
seqdiag {
  APP; service-saksi-management; be-service-member;
  APP -> service-saksi-management [label = "qr_code", note = "GET\n/saksi/validation"]
  service-saksi-management => be-service-member [note = "gRPC\nGetDetailMemberByNo", color = green, label = GetDetailMemberByNoRequest, return = GetDetailMemberByNoResponse]
  APP <-- service-saksi-management [label = "id, member_no, nik, full_name, phone_number, address, dprt_coordinator"]
  APP => service-saksi-management [note = "GET\n/saksi/validation/selfie/{id}.jpg"]
  APP => service-saksi-management [note = "GET\n/saksi/validation/ktp/{id}.jpg"]
  === ACCEPT ===
  APP => service-saksi-management [note = "POST\n/saksi", label = "candidate_id, tps_no"]
  === REJECT ===
  APP => service-saksi-management [note = "DELETE\n/saksi/validation/{id}"]
}
```

5. DPRT accept request register Saksi

```seqdiag
seqdiag {
  APP; service-saksi-management; DB; storage;
  APP -> service-saksi-management [note = "POST\n/saksi", label = "candidate_id, tps_no"]
  service-saksi-management <- DB [label = "Read saksi_candidate data", color = orange]
  service-saksi-management -> DB [label = "Insert data to saksi", color = orange]
  service-saksi-management <- DB [label = "Get saksi_dprt ID", color = orange]
  service-saksi-management -> storage [label = "Safe selfie & KTP photos", color = purple]
  APP <-- service-saksi-management [label = "id, member_no, tps_no"]
}
```

6. DPRT reject request register Saksi

```seqdiag
seqdiag {
  APP; service-saksi-management; DB; storage;
  APP -> service-saksi-management [note = "DELETE\n/saksi/validation/{id}"]
  service-saksi-management -> DB [label = "Delete saksi_candidate data", color = orange]
  service-saksi-management -> storage [label = "Delete selfie & KTP photos", color = purple]
  APP <-- service-saksi-management
}
``` -->

### Entity Relationship Diagram (ERD)

```erd
[recruiter] {bgcolor: "#d0d080"}
  *id {label: "int(11), auto_increment"}
  member_no {label: "varchar(16), not null"}
  *prov_code {label: "int(11), not null"}
  *kab_code {label: "int(11), not null"}
  *kec_code {label: "int(11), not null"}
  kel_code {label: "JSON"}
  *dtm_crt {label: "datetime, not null"}
  *dtm_upd {label: "datetime, not null"}
  log {label: "JSON"}

[saksi] {bgcolor: "#d0f0d0"}
  *id {label: "bigint(20), auto_increment"}
  member_no {label: "varchar(16), not null"}
  +recruiter_id {label: "int(11), not null"}
  *prov_code {label: "int(11), not null"}
  *kab_code {label: "int(11), not null"}
  *kec_code {label: "int(11), not null"}
  *kel_code {label: "int(11), not null"}
  +tps_id_request {label: "bigint(20), not null"}
  +tps_id {label: "bigint(11)"}
  *status {label: "varchar(20), not null, {requested, rejected, waiting, verify, active}"}
  *dtm_crt {label: "datetime, not null"}
  *dtm_upd {label: "datetime, not null"}

[tps] {bgcolor: "#f0b0b0"}
  *id {label: "bigint(20), auto_increment"}
  *prov_code {label: "int(11), not null"}
  *kab_code {label: "int(11), not null"}
  *kec_code {label: "int(11), not null"}
  *kel_code {label: "int(11), not null"}
  *tps_no {label: "int(11), not null"}
  *dtm_crt {label: "datetime, not null"}
  *dtm_upd {label: "datetime, not null"}

[attendance] {bgcolor: "#b0d0f0"}
  *id {label: "bigint(20), auto_increment"}
  +saksi_dprt_id {label: "bigint(20), not null"}
  latitude {label: "number, not null"}
  longitude {label: "number, not null"}
  *dtm_crt {label: "datetime, not null"}

[document_evidence] {bgcolor: "#f0f0a0"}
  *id {label: "bigint(20), auto_increment"}
  +saksi_dprt_id {label: "bigint(20), not null"}
  *document {label: "varchar(255), not null"}
  *compensation {label: "tinyint(1), not null"}
  refused_compensation_reason {label: "text"}
  *dtm_crt {label: "datetime, not null"}

[tutorial] {bgcolor: "#f0b0f0"}
  *id {label: "int(11), auto_increment"}
  *name {label: "varchar(255), not null"}
  *role {label: "varchar(6), not null"}
  *type {label: "varchar(5), not null"}
  category {label: "varchar(6)"}
  *url {label: "varchar(1024), not null"}

recruiter 1--* saksi {label: "id ---> recruiter_id"}
saksi *--1 tps {label: "tps_id_recruiter <--- id"}
saksi *--1 tps {label: "tps_id <--- id"}
saksi 1--1 attendance {label: "id ---> saksi_dprt_id"}
saksi 1--1 document_evidence {label: "id ---> saksi_dprt_id"}
```
