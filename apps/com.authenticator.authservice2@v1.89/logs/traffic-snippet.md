TOTP authenticator uses a hard-coded, static string (`TotpAuthenticator`) to
generate the encryption key, resulting in the same key being used for every
piece of information. This, along with a static initialization vector, results
in no randomness being added to the cryptographic process, meaning that repeated
pieces of information can be trivially detected and read.

```
====== BEGIN REQUEST ======
PUT www.googleapis.com/upload/drive/v3/files?
Content-Type: text/plain

kINkyu035fB5RhSlPwiFGJspdlORIjgKxXG/aCLcnrUPDYvDTZWxj0daLko7hKVx2qOcdGFWjE+x/+LaIAagsdrMsJp7HjL+LvQkuKtKGykME1sVkPCYwat1Y1mwr28zlGTSW/vl6Zku+5o2k3/als9O80LowQO/stmta3G+CEcI4kQSfDnBcUctdhhoWiDRigyHrpLT2IPPMQu+P3qFLGfloWM13rULBTKJZm82Qfp16/SvWxnlBFeUFSxfFB6frAd4ajvU+oUH6rMuacczayNe0JjupiJcKoVgX7DUnyrC0yaXJxUQzRJbEFwHzlok7nFEULgEMoxYTw11N1JZsfNtHcq4zGKVGgBSTBpxnEbNA7QLA/QZ3ggfroz3+29NZ4BVKkc+a/7QQg9k8wi1kmDXPpHu0s8nO5dQZANCCiWjzJzv+ktQx5GDmX+BGvTPQjju0jCY+coasQKw2a/DGn7CQxtsR5XVISO0Saoud+SVPkBEmmnpcKBSHEIljgvAkbjz7rX2QvIj5iHeZ5qA6Uz687LgONRdc2aCAWFY7TnqMgfDY4O4xaJS7FB+XpdTk0+1zCOaG18e8x7LxMpjaahQVq2SkItGlMM8+Y1alQcd3VBbJqT/qaAMJtF5mVPyVT2mloJBnHbCN4u2bGzhOiSCnTSeLbcWF53BRPz+NpadGw5tWSoKDQnMFTj81NQ0Kb2cR1RrnnR5q1IdHSi37ZGhaQ4Ipfrd+1kx9yRPJ+DJyX1OHD1Ko09t85v9m1dLHXfNMDzLPP7YwUFrtbeHEvhA3Ku7hFkp9tsLvDqAsrhEbe8F6vOs/5LvTmo0RDI/mze6bObTgTSjDDz5eauo433DjekFlZnjUIi0QpgYyULKioZKPAFN8COZZzs5dy1lJkVNfVc6JcQQNJU/HJVFyzDe+U21/kyzPok0pMXZtqV94vAujRgTUowBeFbbhBvV0GmUDEVHbeVDXJt7ieZa77QlVfgaPRRyJ2E74ExhL4M=
====== END REQUEST ======
```
