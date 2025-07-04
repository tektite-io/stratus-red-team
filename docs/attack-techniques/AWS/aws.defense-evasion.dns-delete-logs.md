---
title: Delete DNS query logs
---

# Delete DNS query logs




Platform: AWS

## Mappings

- MITRE ATT&CK
    - Defense Evasion



## Description


Deletes a Route53 DNS Resolver query logging configuration. Simulates an attacker disrupting DNS logging.

<span style="font-variant: small-caps;">Warm-up</span>:

- Create a DNS logging configuration.

<span style="font-variant: small-caps;">Detonation</span>:

- Delete the DNS logging configuration using <code>route53:DeleteResolverQueryLogConfig</code>.

## Instructions

```bash title="Detonate with Stratus Red Team"
stratus detonate aws.defense-evasion.dns-delete-logs
```
## Detection


Identify when a DNS logging configuration is deleted, through CloudTrail's <code>DeleteResolverQueryLogConfig</code> event.



## Detonation logs <span class="smallcaps w3-badge w3-light-green w3-round w3-text-sand">new!</span>

The following CloudTrail events are generated when this technique is detonated[^1]:


- `route53resolver:DeleteResolverQueryLogConfig`


??? "View raw detonation logs"

    ```json hl_lines="6"

    [
	   {
	      "awsRegion": "sa-central-3r",
	      "eventCategory": "Management",
	      "eventID": "ba4609ca-b420-4cb6-bdff-307729b3b7db",
	      "eventName": "DeleteResolverQueryLogConfig",
	      "eventSource": "route53resolver.amazonaws.com",
	      "eventTime": "2024-07-31T14:23:46Z",
	      "eventType": "AwsApiCall",
	      "eventVersion": "1.08",
	      "managementEvent": true,
	      "readOnly": false,
	      "recipientAccountId": "206821776919",
	      "requestID": "6dbefe3c-b575-499a-a94d-a3bda0e4009a",
	      "requestParameters": {
	         "originSequenceNumber": 0,
	         "resolverQueryLogConfigId": "rqlc-4473f20ca554c07"
	      },
	      "responseElements": {
	         "resolverQueryLogConfig": {
	            "arn": "arn:aws:route53resolver:sa-central-3r:206821776919:resolver-query-log-config/rqlc-4473f20ca554c07",
	            "associationCount": 0,
	            "creationTime": "2024-07-31T14:23:44.841442289Z",
	            "creatorRequestId": "tf-r53-resolver-query-log-config-20240731142344425800000001",
	            "destinationArn": "arn:aws:s3:::stratus-red-team-dns-delete-bucket-bxxclslsdp",
	            "id": "rqlc-4473f20ca554c07",
	            "name": "stratus-red-team-dns-delete-config-bxxclslsdp",
	            "ownerId": "206821776919",
	            "shareStatus": "NOT_SHARED",
	            "status": "DELETING"
	         }
	      },
	      "sourceIPAddress": "251.234.045.249",
	      "tlsDetails": {
	         "cipherSuite": "TLS_AES_128_GCM_SHA256",
	         "clientProvidedHostHeader": "route53resolver.sa-central-3r.amazonaws.com",
	         "tlsVersion": "TLSv1.3"
	      },
	      "userAgent": "stratus-red-team_bdd216cd-7fb9-4b18-971a-cb585947fd95",
	      "userIdentity": {
	         "accessKeyId": "AKIADT99GZBZR7NVDT0D",
	         "accountId": "206821776919",
	         "arn": "arn:aws:iam::206821776919:user/christophe",
	         "principalId": "AIDAKUK081EB3L71EAZV",
	         "type": "IAMUser",
	         "userName": "christophe"
	      }
	   }
	]
    ```

[^1]: These logs have been gathered from a real detonation of this technique in a test environment using [Grimoire](https://github.com/DataDog/grimoire), and anonymized using [LogLicker](https://github.com/Permiso-io-tools/LogLicker).
