---
title: Usage of ssm:SendCommand on multiple instances
---

# Usage of ssm:SendCommand on multiple instances

 <span class="smallcaps w3-badge w3-orange w3-round w3-text-sand" title="This attack technique might be slow to warm up or detonate">slow</span> 
 <span class="smallcaps w3-badge w3-blue w3-round w3-text-white" title="This attack technique can be detonated multiple times">idempotent</span> 

Platform: AWS

## Mappings

- MITRE ATT&CK
    - Execution



## Description


Simulates an attacker utilizing AWS Systems Manager (SSM) to execute commands through SendCommand on multiple EC2 instances.

<span style="font-variant: small-caps;">Warm-up</span>:

- Create multiple EC2 instances and a VPC (takes a few minutes).

<span style="font-variant: small-caps;">Detonation</span>: 

- Runs <code>ssm:SendCommand</code> on several EC2 instances, to execute the command <code>echo "id=$(id), hostname=$(hostname)"</code> on each of them.

References:

- https://hackingthe.cloud/aws/post_exploitation/run_shell_commands_on_ec2/#send-command
- https://www.chrisfarris.com/post/aws-ir/
- https://www.invictus-ir.com/news/aws-cloudtrail-cheat-sheet
- https://securitycafe.ro/2023/01/17/aws-post-explitation-with-ssm-sendcommand/


## Instructions

```bash title="Detonate with Stratus Red Team"
stratus detonate aws.execution.ssm-send-command
```
## Detection


Identify, through CloudTrail's <code>SendCommand</code> event, especially when <code>requestParameters.instanceIds</code> contains several instances. Sample event:

```json
{
  "eventSource": "ssm.amazonaws.com",
  "eventName": "SendCommand",
  "requestParameters": {
    "instanceIds": [
      "i-0f364762ca43f9661",
      "i-0a86d1f61db2b9b5d",
      "i-08a69bfbe21c67e70"
    ],
    "documentName": "AWS-RunShellScript",
    "parameters": "HIDDEN_DUE_TO_SECURITY_REASONS",
    "interactive": false
  }
}
```

While this technique uses a single call to <code>ssm:SendCommand</code> on several instances, an attacker may use one call per instance to execute commands on. In that case, the <code>SendCommand</code> event will be emitted for each call.



## Detonation logs <span class="smallcaps w3-badge w3-light-green w3-round w3-text-sand">new!</span>

The following CloudTrail events are generated when this technique is detonated[^1]:


- `ssm:DescribeInstanceInformation`

- `ssm:GetCommandInvocation`

- `ssm:SendCommand`


??? "View raw detonation logs"

    ```json hl_lines="6 48 90 132 174 216 258 300 342 384 426 468 510 552 594 636 678 720 762 804 846 888 930 972 1006 1040 1074 1162 1204 1238 1280 1322 1364 1406 1448 1490 1532 1574 1616 1658 1700 1742 1784 1826 1868 1910 1952 1994 2036 2078 2120 2162 2204 2246 2288 2330"

    [
	   {
	      "awsRegion": "sagov-westsouth-1r",
	      "eventCategory": "Management",
	      "eventID": "4723aee9-d1e5-4e32-b48c-0ec39a6d84ea",
	      "eventName": "DescribeInstanceInformation",
	      "eventSource": "ssm.amazonaws.com",
	      "eventTime": "2024-08-02T09:08:27Z",
	      "eventType": "AwsApiCall",
	      "eventVersion": "1.08",
	      "managementEvent": true,
	      "readOnly": true,
	      "recipientAccountId": "056392974792",
	      "requestID": "6edac2c5-52c8-4de5-9d8f-2d1bdc2f9e8b",
	      "requestParameters": {
	         "filters": [
	            {
	               "key": "InstanceIds",
	               "values": [
	                  "i-00456A8D163f546Ff",
	                  "i-cfE23b1a7ceba6f86",
	                  "i-9D40CCFc0aE91CFa5"
	               ]
	            }
	         ]
	      },
	      "responseElements": null,
	      "sourceIPAddress": "253.252.51.07",
	      "tlsDetails": {
	         "cipherSuite": "ECDHE-RSA-AES128-GCM-SHA256",
	         "clientProvidedHostHeader": "ssm.sagov-westsouth-1r.amazonaws.com",
	         "tlsVersion": "TLSv1.2"
	      },
	      "userAgent": "stratus-red-team_ea782787-a65d-4fc4-9fca-1c97869a9a25",
	      "userIdentity": {
	         "accessKeyId": "AKIAW9X2Q2U25SK79UCX",
	         "accountId": "056392974792",
	         "arn": "arn:aws:iam::056392974792:user/christophe",
	         "principalId": "AIDA10CZIPPG73T21TDI",
	         "type": "IAMUser",
	         "userName": "christophe"
	      }
	   },
	   {
	      "awsRegion": "sagov-westsouth-1r",
	      "eventCategory": "Management",
	      "eventID": "bbef7fa1-ec6b-42ca-ae50-a95610fc81d3",
	      "eventName": "DescribeInstanceInformation",
	      "eventSource": "ssm.amazonaws.com",
	      "eventTime": "2024-08-02T09:08:26Z",
	      "eventType": "AwsApiCall",
	      "eventVersion": "1.08",
	      "managementEvent": true,
	      "readOnly": true,
	      "recipientAccountId": "056392974792",
	      "requestID": "083a9fde-def5-4328-bbab-1bd8b0c137cb",
	      "requestParameters": {
	         "filters": [
	            {
	               "key": "InstanceIds",
	               "values": [
	                  "i-00456A8D163f546Ff",
	                  "i-cfE23b1a7ceba6f86",
	                  "i-9D40CCFc0aE91CFa5"
	               ]
	            }
	         ]
	      },
	      "responseElements": null,
	      "sourceIPAddress": "253.252.51.07",
	      "tlsDetails": {
	         "cipherSuite": "ECDHE-RSA-AES128-GCM-SHA256",
	         "clientProvidedHostHeader": "ssm.sagov-westsouth-1r.amazonaws.com",
	         "tlsVersion": "TLSv1.2"
	      },
	      "userAgent": "stratus-red-team_ea782787-a65d-4fc4-9fca-1c97869a9a25",
	      "userIdentity": {
	         "accessKeyId": "AKIAW9X2Q2U25SK79UCX",
	         "accountId": "056392974792",
	         "arn": "arn:aws:iam::056392974792:user/christophe",
	         "principalId": "AIDA10CZIPPG73T21TDI",
	         "type": "IAMUser",
	         "userName": "christophe"
	      }
	   },
	   {
	      "awsRegion": "sagov-westsouth-1r",
	      "eventCategory": "Management",
	      "eventID": "d6738500-de0a-4a7d-af41-c42225b1d627",
	      "eventName": "DescribeInstanceInformation",
	      "eventSource": "ssm.amazonaws.com",
	      "eventTime": "2024-08-02T09:08:23Z",
	      "eventType": "AwsApiCall",
	      "eventVersion": "1.08",
	      "managementEvent": true,
	      "readOnly": true,
	      "recipientAccountId": "056392974792",
	      "requestID": "cdf0af8d-32e8-4094-b5ad-0ad6aa898a2b",
	      "requestParameters": {
	         "filters": [
	            {
	               "key": "InstanceIds",
	               "values": [
	                  "i-00456A8D163f546Ff",
	                  "i-cfE23b1a7ceba6f86",
	                  "i-9D40CCFc0aE91CFa5"
	               ]
	            }
	         ]
	      },
	      "responseElements": null,
	      "sourceIPAddress": "253.252.51.07",
	      "tlsDetails": {
	         "cipherSuite": "ECDHE-RSA-AES128-GCM-SHA256",
	         "clientProvidedHostHeader": "ssm.sagov-westsouth-1r.amazonaws.com",
	         "tlsVersion": "TLSv1.2"
	      },
	      "userAgent": "stratus-red-team_ea782787-a65d-4fc4-9fca-1c97869a9a25",
	      "userIdentity": {
	         "accessKeyId": "AKIAW9X2Q2U25SK79UCX",
	         "accountId": "056392974792",
	         "arn": "arn:aws:iam::056392974792:user/christophe",
	         "principalId": "AIDA10CZIPPG73T21TDI",
	         "type": "IAMUser",
	         "userName": "christophe"
	      }
	   },
	   {
	      "awsRegion": "sagov-westsouth-1r",
	      "eventCategory": "Management",
	      "eventID": "5ceab743-d517-46d5-b162-bf881ae0be0c",
	      "eventName": "DescribeInstanceInformation",
	      "eventSource": "ssm.amazonaws.com",
	      "eventTime": "2024-08-02T09:08:21Z",
	      "eventType": "AwsApiCall",
	      "eventVersion": "1.08",
	      "managementEvent": true,
	      "readOnly": true,
	      "recipientAccountId": "056392974792",
	      "requestID": "b48c0a2a-5c9b-4bd9-9e2a-74c84a55aefe",
	      "requestParameters": {
	         "filters": [
	            {
	               "key": "InstanceIds",
	               "values": [
	                  "i-00456A8D163f546Ff",
	                  "i-cfE23b1a7ceba6f86",
	                  "i-9D40CCFc0aE91CFa5"
	               ]
	            }
	         ]
	      },
	      "responseElements": null,
	      "sourceIPAddress": "253.252.51.07",
	      "tlsDetails": {
	         "cipherSuite": "ECDHE-RSA-AES128-GCM-SHA256",
	         "clientProvidedHostHeader": "ssm.sagov-westsouth-1r.amazonaws.com",
	         "tlsVersion": "TLSv1.2"
	      },
	      "userAgent": "stratus-red-team_ea782787-a65d-4fc4-9fca-1c97869a9a25",
	      "userIdentity": {
	         "accessKeyId": "AKIAW9X2Q2U25SK79UCX",
	         "accountId": "056392974792",
	         "arn": "arn:aws:iam::056392974792:user/christophe",
	         "principalId": "AIDA10CZIPPG73T21TDI",
	         "type": "IAMUser",
	         "userName": "christophe"
	      }
	   },
	   {
	      "awsRegion": "sagov-westsouth-1r",
	      "eventCategory": "Management",
	      "eventID": "df4e2a35-15df-4329-9b51-f260dcefba7b",
	      "eventName": "DescribeInstanceInformation",
	      "eventSource": "ssm.amazonaws.com",
	      "eventTime": "2024-08-02T09:08:19Z",
	      "eventType": "AwsApiCall",
	      "eventVersion": "1.08",
	      "managementEvent": true,
	      "readOnly": true,
	      "recipientAccountId": "056392974792",
	      "requestID": "fe3cc368-5dd9-4629-8db6-966b9b396005",
	      "requestParameters": {
	         "filters": [
	            {
	               "key": "InstanceIds",
	               "values": [
	                  "i-00456A8D163f546Ff",
	                  "i-cfE23b1a7ceba6f86",
	                  "i-9D40CCFc0aE91CFa5"
	               ]
	            }
	         ]
	      },
	      "responseElements": null,
	      "sourceIPAddress": "253.252.51.07",
	      "tlsDetails": {
	         "cipherSuite": "ECDHE-RSA-AES128-GCM-SHA256",
	         "clientProvidedHostHeader": "ssm.sagov-westsouth-1r.amazonaws.com",
	         "tlsVersion": "TLSv1.2"
	      },
	      "userAgent": "stratus-red-team_ea782787-a65d-4fc4-9fca-1c97869a9a25",
	      "userIdentity": {
	         "accessKeyId": "AKIAW9X2Q2U25SK79UCX",
	         "accountId": "056392974792",
	         "arn": "arn:aws:iam::056392974792:user/christophe",
	         "principalId": "AIDA10CZIPPG73T21TDI",
	         "type": "IAMUser",
	         "userName": "christophe"
	      }
	   },
	   {
	      "awsRegion": "sagov-westsouth-1r",
	      "eventCategory": "Management",
	      "eventID": "414a9a7c-01f3-4acc-9b55-bf1f677e3a54",
	      "eventName": "DescribeInstanceInformation",
	      "eventSource": "ssm.amazonaws.com",
	      "eventTime": "2024-08-02T09:08:17Z",
	      "eventType": "AwsApiCall",
	      "eventVersion": "1.08",
	      "managementEvent": true,
	      "readOnly": true,
	      "recipientAccountId": "056392974792",
	      "requestID": "6425b4c5-5688-4d8f-8165-cf0b565cdb72",
	      "requestParameters": {
	         "filters": [
	            {
	               "key": "InstanceIds",
	               "values": [
	                  "i-00456A8D163f546Ff",
	                  "i-cfE23b1a7ceba6f86",
	                  "i-9D40CCFc0aE91CFa5"
	               ]
	            }
	         ]
	      },
	      "responseElements": null,
	      "sourceIPAddress": "253.252.51.07",
	      "tlsDetails": {
	         "cipherSuite": "ECDHE-RSA-AES128-GCM-SHA256",
	         "clientProvidedHostHeader": "ssm.sagov-westsouth-1r.amazonaws.com",
	         "tlsVersion": "TLSv1.2"
	      },
	      "userAgent": "stratus-red-team_ea782787-a65d-4fc4-9fca-1c97869a9a25",
	      "userIdentity": {
	         "accessKeyId": "AKIAW9X2Q2U25SK79UCX",
	         "accountId": "056392974792",
	         "arn": "arn:aws:iam::056392974792:user/christophe",
	         "principalId": "AIDA10CZIPPG73T21TDI",
	         "type": "IAMUser",
	         "userName": "christophe"
	      }
	   },
	   {
	      "awsRegion": "sagov-westsouth-1r",
	      "eventCategory": "Management",
	      "eventID": "2c1e26d1-6685-4640-ba79-81149872d066",
	      "eventName": "DescribeInstanceInformation",
	      "eventSource": "ssm.amazonaws.com",
	      "eventTime": "2024-08-02T09:08:16Z",
	      "eventType": "AwsApiCall",
	      "eventVersion": "1.08",
	      "managementEvent": true,
	      "readOnly": true,
	      "recipientAccountId": "056392974792",
	      "requestID": "0ea54e95-cde4-4aec-9ef3-d28f44594966",
	      "requestParameters": {
	         "filters": [
	            {
	               "key": "InstanceIds",
	               "values": [
	                  "i-00456A8D163f546Ff",
	                  "i-cfE23b1a7ceba6f86",
	                  "i-9D40CCFc0aE91CFa5"
	               ]
	            }
	         ]
	      },
	      "responseElements": null,
	      "sourceIPAddress": "253.252.51.07",
	      "tlsDetails": {
	         "cipherSuite": "ECDHE-RSA-AES128-GCM-SHA256",
	         "clientProvidedHostHeader": "ssm.sagov-westsouth-1r.amazonaws.com",
	         "tlsVersion": "TLSv1.2"
	      },
	      "userAgent": "stratus-red-team_ea782787-a65d-4fc4-9fca-1c97869a9a25",
	      "userIdentity": {
	         "accessKeyId": "AKIAW9X2Q2U25SK79UCX",
	         "accountId": "056392974792",
	         "arn": "arn:aws:iam::056392974792:user/christophe",
	         "principalId": "AIDA10CZIPPG73T21TDI",
	         "type": "IAMUser",
	         "userName": "christophe"
	      }
	   },
	   {
	      "awsRegion": "sagov-westsouth-1r",
	      "eventCategory": "Management",
	      "eventID": "a4ca6ef1-b00e-476a-8dcf-6b1b2e75b335",
	      "eventName": "DescribeInstanceInformation",
	      "eventSource": "ssm.amazonaws.com",
	      "eventTime": "2024-08-02T09:08:15Z",
	      "eventType": "AwsApiCall",
	      "eventVersion": "1.08",
	      "managementEvent": true,
	      "readOnly": true,
	      "recipientAccountId": "056392974792",
	      "requestID": "0c49d64c-5995-485c-930f-fbb3fcda42ab",
	      "requestParameters": {
	         "filters": [
	            {
	               "key": "InstanceIds",
	               "values": [
	                  "i-00456A8D163f546Ff",
	                  "i-cfE23b1a7ceba6f86",
	                  "i-9D40CCFc0aE91CFa5"
	               ]
	            }
	         ]
	      },
	      "responseElements": null,
	      "sourceIPAddress": "253.252.51.07",
	      "tlsDetails": {
	         "cipherSuite": "ECDHE-RSA-AES128-GCM-SHA256",
	         "clientProvidedHostHeader": "ssm.sagov-westsouth-1r.amazonaws.com",
	         "tlsVersion": "TLSv1.2"
	      },
	      "userAgent": "stratus-red-team_ea782787-a65d-4fc4-9fca-1c97869a9a25",
	      "userIdentity": {
	         "accessKeyId": "AKIAW9X2Q2U25SK79UCX",
	         "accountId": "056392974792",
	         "arn": "arn:aws:iam::056392974792:user/christophe",
	         "principalId": "AIDA10CZIPPG73T21TDI",
	         "type": "IAMUser",
	         "userName": "christophe"
	      }
	   },
	   {
	      "awsRegion": "sagov-westsouth-1r",
	      "eventCategory": "Management",
	      "eventID": "2b3aacaa-3e89-405c-b53b-f99a0555661d",
	      "eventName": "DescribeInstanceInformation",
	      "eventSource": "ssm.amazonaws.com",
	      "eventTime": "2024-08-02T09:08:14Z",
	      "eventType": "AwsApiCall",
	      "eventVersion": "1.08",
	      "managementEvent": true,
	      "readOnly": true,
	      "recipientAccountId": "056392974792",
	      "requestID": "2abe2e44-53f2-4207-825e-dc569c2be9f5",
	      "requestParameters": {
	         "filters": [
	            {
	               "key": "InstanceIds",
	               "values": [
	                  "i-00456A8D163f546Ff",
	                  "i-cfE23b1a7ceba6f86",
	                  "i-9D40CCFc0aE91CFa5"
	               ]
	            }
	         ]
	      },
	      "responseElements": null,
	      "sourceIPAddress": "253.252.51.07",
	      "tlsDetails": {
	         "cipherSuite": "ECDHE-RSA-AES128-GCM-SHA256",
	         "clientProvidedHostHeader": "ssm.sagov-westsouth-1r.amazonaws.com",
	         "tlsVersion": "TLSv1.2"
	      },
	      "userAgent": "stratus-red-team_ea782787-a65d-4fc4-9fca-1c97869a9a25",
	      "userIdentity": {
	         "accessKeyId": "AKIAW9X2Q2U25SK79UCX",
	         "accountId": "056392974792",
	         "arn": "arn:aws:iam::056392974792:user/christophe",
	         "principalId": "AIDA10CZIPPG73T21TDI",
	         "type": "IAMUser",
	         "userName": "christophe"
	      }
	   },
	   {
	      "awsRegion": "sagov-westsouth-1r",
	      "eventCategory": "Management",
	      "eventID": "2610da37-3b46-48b2-82b3-59e0c77c9db0",
	      "eventName": "DescribeInstanceInformation",
	      "eventSource": "ssm.amazonaws.com",
	      "eventTime": "2024-08-02T09:08:13Z",
	      "eventType": "AwsApiCall",
	      "eventVersion": "1.08",
	      "managementEvent": true,
	      "readOnly": true,
	      "recipientAccountId": "056392974792",
	      "requestID": "c2320169-a590-4aa4-bfbe-73d0eef783fa",
	      "requestParameters": {
	         "filters": [
	            {
	               "key": "InstanceIds",
	               "values": [
	                  "i-00456A8D163f546Ff",
	                  "i-cfE23b1a7ceba6f86",
	                  "i-9D40CCFc0aE91CFa5"
	               ]
	            }
	         ]
	      },
	      "responseElements": null,
	      "sourceIPAddress": "253.252.51.07",
	      "tlsDetails": {
	         "cipherSuite": "ECDHE-RSA-AES128-GCM-SHA256",
	         "clientProvidedHostHeader": "ssm.sagov-westsouth-1r.amazonaws.com",
	         "tlsVersion": "TLSv1.2"
	      },
	      "userAgent": "stratus-red-team_ea782787-a65d-4fc4-9fca-1c97869a9a25",
	      "userIdentity": {
	         "accessKeyId": "AKIAW9X2Q2U25SK79UCX",
	         "accountId": "056392974792",
	         "arn": "arn:aws:iam::056392974792:user/christophe",
	         "principalId": "AIDA10CZIPPG73T21TDI",
	         "type": "IAMUser",
	         "userName": "christophe"
	      }
	   },
	   {
	      "awsRegion": "sagov-westsouth-1r",
	      "eventCategory": "Management",
	      "eventID": "04151503-f5e2-4356-abdd-14b08e2285ef",
	      "eventName": "DescribeInstanceInformation",
	      "eventSource": "ssm.amazonaws.com",
	      "eventTime": "2024-08-02T09:08:12Z",
	      "eventType": "AwsApiCall",
	      "eventVersion": "1.08",
	      "managementEvent": true,
	      "readOnly": true,
	      "recipientAccountId": "056392974792",
	      "requestID": "61a85904-a3b8-4dd6-aaef-2efd548cf9ae",
	      "requestParameters": {
	         "filters": [
	            {
	               "key": "InstanceIds",
	               "values": [
	                  "i-00456A8D163f546Ff",
	                  "i-cfE23b1a7ceba6f86",
	                  "i-9D40CCFc0aE91CFa5"
	               ]
	            }
	         ]
	      },
	      "responseElements": null,
	      "sourceIPAddress": "253.252.51.07",
	      "tlsDetails": {
	         "cipherSuite": "ECDHE-RSA-AES128-GCM-SHA256",
	         "clientProvidedHostHeader": "ssm.sagov-westsouth-1r.amazonaws.com",
	         "tlsVersion": "TLSv1.2"
	      },
	      "userAgent": "stratus-red-team_ea782787-a65d-4fc4-9fca-1c97869a9a25",
	      "userIdentity": {
	         "accessKeyId": "AKIAW9X2Q2U25SK79UCX",
	         "accountId": "056392974792",
	         "arn": "arn:aws:iam::056392974792:user/christophe",
	         "principalId": "AIDA10CZIPPG73T21TDI",
	         "type": "IAMUser",
	         "userName": "christophe"
	      }
	   },
	   {
	      "awsRegion": "sagov-westsouth-1r",
	      "eventCategory": "Management",
	      "eventID": "898fc3e2-242e-48f1-a560-8b835d90bdee",
	      "eventName": "DescribeInstanceInformation",
	      "eventSource": "ssm.amazonaws.com",
	      "eventTime": "2024-08-02T09:08:10Z",
	      "eventType": "AwsApiCall",
	      "eventVersion": "1.08",
	      "managementEvent": true,
	      "readOnly": true,
	      "recipientAccountId": "056392974792",
	      "requestID": "8931849b-3dbb-440f-ac27-1fb5d4890d3b",
	      "requestParameters": {
	         "filters": [
	            {
	               "key": "InstanceIds",
	               "values": [
	                  "i-00456A8D163f546Ff",
	                  "i-cfE23b1a7ceba6f86",
	                  "i-9D40CCFc0aE91CFa5"
	               ]
	            }
	         ]
	      },
	      "responseElements": null,
	      "sourceIPAddress": "253.252.51.07",
	      "tlsDetails": {
	         "cipherSuite": "ECDHE-RSA-AES128-GCM-SHA256",
	         "clientProvidedHostHeader": "ssm.sagov-westsouth-1r.amazonaws.com",
	         "tlsVersion": "TLSv1.2"
	      },
	      "userAgent": "stratus-red-team_ea782787-a65d-4fc4-9fca-1c97869a9a25",
	      "userIdentity": {
	         "accessKeyId": "AKIAW9X2Q2U25SK79UCX",
	         "accountId": "056392974792",
	         "arn": "arn:aws:iam::056392974792:user/christophe",
	         "principalId": "AIDA10CZIPPG73T21TDI",
	         "type": "IAMUser",
	         "userName": "christophe"
	      }
	   },
	   {
	      "awsRegion": "sagov-westsouth-1r",
	      "eventCategory": "Management",
	      "eventID": "411687aa-d840-40f7-ae31-adb0619c0401",
	      "eventName": "DescribeInstanceInformation",
	      "eventSource": "ssm.amazonaws.com",
	      "eventTime": "2024-08-02T09:08:09Z",
	      "eventType": "AwsApiCall",
	      "eventVersion": "1.08",
	      "managementEvent": true,
	      "readOnly": true,
	      "recipientAccountId": "056392974792",
	      "requestID": "003bfa5a-ef20-46b7-bf79-8a11a49ab14e",
	      "requestParameters": {
	         "filters": [
	            {
	               "key": "InstanceIds",
	               "values": [
	                  "i-00456A8D163f546Ff",
	                  "i-cfE23b1a7ceba6f86",
	                  "i-9D40CCFc0aE91CFa5"
	               ]
	            }
	         ]
	      },
	      "responseElements": null,
	      "sourceIPAddress": "253.252.51.07",
	      "tlsDetails": {
	         "cipherSuite": "ECDHE-RSA-AES128-GCM-SHA256",
	         "clientProvidedHostHeader": "ssm.sagov-westsouth-1r.amazonaws.com",
	         "tlsVersion": "TLSv1.2"
	      },
	      "userAgent": "stratus-red-team_ea782787-a65d-4fc4-9fca-1c97869a9a25",
	      "userIdentity": {
	         "accessKeyId": "AKIAW9X2Q2U25SK79UCX",
	         "accountId": "056392974792",
	         "arn": "arn:aws:iam::056392974792:user/christophe",
	         "principalId": "AIDA10CZIPPG73T21TDI",
	         "type": "IAMUser",
	         "userName": "christophe"
	      }
	   },
	   {
	      "awsRegion": "sagov-westsouth-1r",
	      "eventCategory": "Management",
	      "eventID": "ff20ced4-0e3c-42a7-9ed9-f32cd2cbb672",
	      "eventName": "DescribeInstanceInformation",
	      "eventSource": "ssm.amazonaws.com",
	      "eventTime": "2024-08-02T09:08:08Z",
	      "eventType": "AwsApiCall",
	      "eventVersion": "1.08",
	      "managementEvent": true,
	      "readOnly": true,
	      "recipientAccountId": "056392974792",
	      "requestID": "0234c68e-9ebe-4fc5-81ab-798de9bdc451",
	      "requestParameters": {
	         "filters": [
	            {
	               "key": "InstanceIds",
	               "values": [
	                  "i-00456A8D163f546Ff",
	                  "i-cfE23b1a7ceba6f86",
	                  "i-9D40CCFc0aE91CFa5"
	               ]
	            }
	         ]
	      },
	      "responseElements": null,
	      "sourceIPAddress": "253.252.51.07",
	      "tlsDetails": {
	         "cipherSuite": "ECDHE-RSA-AES128-GCM-SHA256",
	         "clientProvidedHostHeader": "ssm.sagov-westsouth-1r.amazonaws.com",
	         "tlsVersion": "TLSv1.2"
	      },
	      "userAgent": "stratus-red-team_ea782787-a65d-4fc4-9fca-1c97869a9a25",
	      "userIdentity": {
	         "accessKeyId": "AKIAW9X2Q2U25SK79UCX",
	         "accountId": "056392974792",
	         "arn": "arn:aws:iam::056392974792:user/christophe",
	         "principalId": "AIDA10CZIPPG73T21TDI",
	         "type": "IAMUser",
	         "userName": "christophe"
	      }
	   },
	   {
	      "awsRegion": "sagov-westsouth-1r",
	      "eventCategory": "Management",
	      "eventID": "e0643796-b464-4e13-8680-00c6dc57ef72",
	      "eventName": "DescribeInstanceInformation",
	      "eventSource": "ssm.amazonaws.com",
	      "eventTime": "2024-08-02T09:08:07Z",
	      "eventType": "AwsApiCall",
	      "eventVersion": "1.08",
	      "managementEvent": true,
	      "readOnly": true,
	      "recipientAccountId": "056392974792",
	      "requestID": "1543ba41-1625-45c3-8f4f-ab5463d68b02",
	      "requestParameters": {
	         "filters": [
	            {
	               "key": "InstanceIds",
	               "values": [
	                  "i-00456A8D163f546Ff",
	                  "i-cfE23b1a7ceba6f86",
	                  "i-9D40CCFc0aE91CFa5"
	               ]
	            }
	         ]
	      },
	      "responseElements": null,
	      "sourceIPAddress": "253.252.51.07",
	      "tlsDetails": {
	         "cipherSuite": "ECDHE-RSA-AES128-GCM-SHA256",
	         "clientProvidedHostHeader": "ssm.sagov-westsouth-1r.amazonaws.com",
	         "tlsVersion": "TLSv1.2"
	      },
	      "userAgent": "stratus-red-team_ea782787-a65d-4fc4-9fca-1c97869a9a25",
	      "userIdentity": {
	         "accessKeyId": "AKIAW9X2Q2U25SK79UCX",
	         "accountId": "056392974792",
	         "arn": "arn:aws:iam::056392974792:user/christophe",
	         "principalId": "AIDA10CZIPPG73T21TDI",
	         "type": "IAMUser",
	         "userName": "christophe"
	      }
	   },
	   {
	      "awsRegion": "sagov-westsouth-1r",
	      "eventCategory": "Management",
	      "eventID": "1540ea9a-4d6b-45b5-b84d-e9711e7801fb",
	      "eventName": "DescribeInstanceInformation",
	      "eventSource": "ssm.amazonaws.com",
	      "eventTime": "2024-08-02T09:08:06Z",
	      "eventType": "AwsApiCall",
	      "eventVersion": "1.08",
	      "managementEvent": true,
	      "readOnly": true,
	      "recipientAccountId": "056392974792",
	      "requestID": "0e53ee03-5e82-4bcc-80fe-1f5929260121",
	      "requestParameters": {
	         "filters": [
	            {
	               "key": "InstanceIds",
	               "values": [
	                  "i-00456A8D163f546Ff",
	                  "i-cfE23b1a7ceba6f86",
	                  "i-9D40CCFc0aE91CFa5"
	               ]
	            }
	         ]
	      },
	      "responseElements": null,
	      "sourceIPAddress": "253.252.51.07",
	      "tlsDetails": {
	         "cipherSuite": "ECDHE-RSA-AES128-GCM-SHA256",
	         "clientProvidedHostHeader": "ssm.sagov-westsouth-1r.amazonaws.com",
	         "tlsVersion": "TLSv1.2"
	      },
	      "userAgent": "stratus-red-team_ea782787-a65d-4fc4-9fca-1c97869a9a25",
	      "userIdentity": {
	         "accessKeyId": "AKIAW9X2Q2U25SK79UCX",
	         "accountId": "056392974792",
	         "arn": "arn:aws:iam::056392974792:user/christophe",
	         "principalId": "AIDA10CZIPPG73T21TDI",
	         "type": "IAMUser",
	         "userName": "christophe"
	      }
	   },
	   {
	      "awsRegion": "sagov-westsouth-1r",
	      "eventCategory": "Management",
	      "eventID": "0d989ab9-09ae-44c4-9dc8-3f3c9aa4f4b1",
	      "eventName": "DescribeInstanceInformation",
	      "eventSource": "ssm.amazonaws.com",
	      "eventTime": "2024-08-02T09:08:05Z",
	      "eventType": "AwsApiCall",
	      "eventVersion": "1.08",
	      "managementEvent": true,
	      "readOnly": true,
	      "recipientAccountId": "056392974792",
	      "requestID": "cdea3227-f206-4316-8ba4-980b36f6124a",
	      "requestParameters": {
	         "filters": [
	            {
	               "key": "InstanceIds",
	               "values": [
	                  "i-00456A8D163f546Ff",
	                  "i-cfE23b1a7ceba6f86",
	                  "i-9D40CCFc0aE91CFa5"
	               ]
	            }
	         ]
	      },
	      "responseElements": null,
	      "sourceIPAddress": "253.252.51.07",
	      "tlsDetails": {
	         "cipherSuite": "ECDHE-RSA-AES128-GCM-SHA256",
	         "clientProvidedHostHeader": "ssm.sagov-westsouth-1r.amazonaws.com",
	         "tlsVersion": "TLSv1.2"
	      },
	      "userAgent": "stratus-red-team_ea782787-a65d-4fc4-9fca-1c97869a9a25",
	      "userIdentity": {
	         "accessKeyId": "AKIAW9X2Q2U25SK79UCX",
	         "accountId": "056392974792",
	         "arn": "arn:aws:iam::056392974792:user/christophe",
	         "principalId": "AIDA10CZIPPG73T21TDI",
	         "type": "IAMUser",
	         "userName": "christophe"
	      }
	   },
	   {
	      "awsRegion": "sagov-westsouth-1r",
	      "eventCategory": "Management",
	      "eventID": "ab4521b5-0b95-4e01-bc57-9124138b6d07",
	      "eventName": "DescribeInstanceInformation",
	      "eventSource": "ssm.amazonaws.com",
	      "eventTime": "2024-08-02T09:08:04Z",
	      "eventType": "AwsApiCall",
	      "eventVersion": "1.08",
	      "managementEvent": true,
	      "readOnly": true,
	      "recipientAccountId": "056392974792",
	      "requestID": "581d7a02-356c-4b34-88ff-0570f6fb1d2b",
	      "requestParameters": {
	         "filters": [
	            {
	               "key": "InstanceIds",
	               "values": [
	                  "i-00456A8D163f546Ff",
	                  "i-cfE23b1a7ceba6f86",
	                  "i-9D40CCFc0aE91CFa5"
	               ]
	            }
	         ]
	      },
	      "responseElements": null,
	      "sourceIPAddress": "253.252.51.07",
	      "tlsDetails": {
	         "cipherSuite": "ECDHE-RSA-AES128-GCM-SHA256",
	         "clientProvidedHostHeader": "ssm.sagov-westsouth-1r.amazonaws.com",
	         "tlsVersion": "TLSv1.2"
	      },
	      "userAgent": "stratus-red-team_ea782787-a65d-4fc4-9fca-1c97869a9a25",
	      "userIdentity": {
	         "accessKeyId": "AKIAW9X2Q2U25SK79UCX",
	         "accountId": "056392974792",
	         "arn": "arn:aws:iam::056392974792:user/christophe",
	         "principalId": "AIDA10CZIPPG73T21TDI",
	         "type": "IAMUser",
	         "userName": "christophe"
	      }
	   },
	   {
	      "awsRegion": "sagov-westsouth-1r",
	      "eventCategory": "Management",
	      "eventID": "150f7722-557f-47a7-849c-5c44cba78e2e",
	      "eventName": "DescribeInstanceInformation",
	      "eventSource": "ssm.amazonaws.com",
	      "eventTime": "2024-08-02T09:08:02Z",
	      "eventType": "AwsApiCall",
	      "eventVersion": "1.08",
	      "managementEvent": true,
	      "readOnly": true,
	      "recipientAccountId": "056392974792",
	      "requestID": "3674ec77-adc1-4474-aad5-a1a6fed8b8d4",
	      "requestParameters": {
	         "filters": [
	            {
	               "key": "InstanceIds",
	               "values": [
	                  "i-00456A8D163f546Ff",
	                  "i-cfE23b1a7ceba6f86",
	                  "i-9D40CCFc0aE91CFa5"
	               ]
	            }
	         ]
	      },
	      "responseElements": null,
	      "sourceIPAddress": "253.252.51.07",
	      "tlsDetails": {
	         "cipherSuite": "ECDHE-RSA-AES128-GCM-SHA256",
	         "clientProvidedHostHeader": "ssm.sagov-westsouth-1r.amazonaws.com",
	         "tlsVersion": "TLSv1.2"
	      },
	      "userAgent": "stratus-red-team_ea782787-a65d-4fc4-9fca-1c97869a9a25",
	      "userIdentity": {
	         "accessKeyId": "AKIAW9X2Q2U25SK79UCX",
	         "accountId": "056392974792",
	         "arn": "arn:aws:iam::056392974792:user/christophe",
	         "principalId": "AIDA10CZIPPG73T21TDI",
	         "type": "IAMUser",
	         "userName": "christophe"
	      }
	   },
	   {
	      "awsRegion": "sagov-westsouth-1r",
	      "eventCategory": "Management",
	      "eventID": "c68a4a51-cfc2-490d-86da-f0aff1e000e6",
	      "eventName": "DescribeInstanceInformation",
	      "eventSource": "ssm.amazonaws.com",
	      "eventTime": "2024-08-02T09:08:01Z",
	      "eventType": "AwsApiCall",
	      "eventVersion": "1.08",
	      "managementEvent": true,
	      "readOnly": true,
	      "recipientAccountId": "056392974792",
	      "requestID": "ab1a6ced-43d6-459c-b67b-6c1acb255fd8",
	      "requestParameters": {
	         "filters": [
	            {
	               "key": "InstanceIds",
	               "values": [
	                  "i-00456A8D163f546Ff",
	                  "i-cfE23b1a7ceba6f86",
	                  "i-9D40CCFc0aE91CFa5"
	               ]
	            }
	         ]
	      },
	      "responseElements": null,
	      "sourceIPAddress": "253.252.51.07",
	      "tlsDetails": {
	         "cipherSuite": "ECDHE-RSA-AES128-GCM-SHA256",
	         "clientProvidedHostHeader": "ssm.sagov-westsouth-1r.amazonaws.com",
	         "tlsVersion": "TLSv1.2"
	      },
	      "userAgent": "stratus-red-team_ea782787-a65d-4fc4-9fca-1c97869a9a25",
	      "userIdentity": {
	         "accessKeyId": "AKIAW9X2Q2U25SK79UCX",
	         "accountId": "056392974792",
	         "arn": "arn:aws:iam::056392974792:user/christophe",
	         "principalId": "AIDA10CZIPPG73T21TDI",
	         "type": "IAMUser",
	         "userName": "christophe"
	      }
	   },
	   {
	      "awsRegion": "sagov-westsouth-1r",
	      "eventCategory": "Management",
	      "eventID": "2582b47b-76b8-4eb4-a455-9f97b000d38a",
	      "eventName": "DescribeInstanceInformation",
	      "eventSource": "ssm.amazonaws.com",
	      "eventTime": "2024-08-02T09:08:00Z",
	      "eventType": "AwsApiCall",
	      "eventVersion": "1.08",
	      "managementEvent": true,
	      "readOnly": true,
	      "recipientAccountId": "056392974792",
	      "requestID": "fe6366b5-7c41-4a98-ab58-fa895d8d71f8",
	      "requestParameters": {
	         "filters": [
	            {
	               "key": "InstanceIds",
	               "values": [
	                  "i-00456A8D163f546Ff",
	                  "i-cfE23b1a7ceba6f86",
	                  "i-9D40CCFc0aE91CFa5"
	               ]
	            }
	         ]
	      },
	      "responseElements": null,
	      "sourceIPAddress": "253.252.51.07",
	      "tlsDetails": {
	         "cipherSuite": "ECDHE-RSA-AES128-GCM-SHA256",
	         "clientProvidedHostHeader": "ssm.sagov-westsouth-1r.amazonaws.com",
	         "tlsVersion": "TLSv1.2"
	      },
	      "userAgent": "stratus-red-team_ea782787-a65d-4fc4-9fca-1c97869a9a25",
	      "userIdentity": {
	         "accessKeyId": "AKIAW9X2Q2U25SK79UCX",
	         "accountId": "056392974792",
	         "arn": "arn:aws:iam::056392974792:user/christophe",
	         "principalId": "AIDA10CZIPPG73T21TDI",
	         "type": "IAMUser",
	         "userName": "christophe"
	      }
	   },
	   {
	      "awsRegion": "sagov-westsouth-1r",
	      "eventCategory": "Management",
	      "eventID": "aa35aa1c-1989-4beb-a540-2a47b88a2119",
	      "eventName": "DescribeInstanceInformation",
	      "eventSource": "ssm.amazonaws.com",
	      "eventTime": "2024-08-02T09:07:59Z",
	      "eventType": "AwsApiCall",
	      "eventVersion": "1.08",
	      "managementEvent": true,
	      "readOnly": true,
	      "recipientAccountId": "056392974792",
	      "requestID": "7c848a81-1e4b-4457-a067-ede23efb8f96",
	      "requestParameters": {
	         "filters": [
	            {
	               "key": "InstanceIds",
	               "values": [
	                  "i-00456A8D163f546Ff",
	                  "i-cfE23b1a7ceba6f86",
	                  "i-9D40CCFc0aE91CFa5"
	               ]
	            }
	         ]
	      },
	      "responseElements": null,
	      "sourceIPAddress": "253.252.51.07",
	      "tlsDetails": {
	         "cipherSuite": "ECDHE-RSA-AES128-GCM-SHA256",
	         "clientProvidedHostHeader": "ssm.sagov-westsouth-1r.amazonaws.com",
	         "tlsVersion": "TLSv1.2"
	      },
	      "userAgent": "stratus-red-team_ea782787-a65d-4fc4-9fca-1c97869a9a25",
	      "userIdentity": {
	         "accessKeyId": "AKIAW9X2Q2U25SK79UCX",
	         "accountId": "056392974792",
	         "arn": "arn:aws:iam::056392974792:user/christophe",
	         "principalId": "AIDA10CZIPPG73T21TDI",
	         "type": "IAMUser",
	         "userName": "christophe"
	      }
	   },
	   {
	      "awsRegion": "sagov-westsouth-1r",
	      "eventCategory": "Management",
	      "eventID": "0d86f878-d8c0-475c-8079-2a1243666e45",
	      "eventName": "DescribeInstanceInformation",
	      "eventSource": "ssm.amazonaws.com",
	      "eventTime": "2024-08-02T09:07:58Z",
	      "eventType": "AwsApiCall",
	      "eventVersion": "1.08",
	      "managementEvent": true,
	      "readOnly": true,
	      "recipientAccountId": "056392974792",
	      "requestID": "914d4883-5725-4059-bf32-8b240cd2be40",
	      "requestParameters": {
	         "filters": [
	            {
	               "key": "InstanceIds",
	               "values": [
	                  "i-00456A8D163f546Ff",
	                  "i-cfE23b1a7ceba6f86",
	                  "i-9D40CCFc0aE91CFa5"
	               ]
	            }
	         ]
	      },
	      "responseElements": null,
	      "sourceIPAddress": "253.252.51.07",
	      "tlsDetails": {
	         "cipherSuite": "ECDHE-RSA-AES128-GCM-SHA256",
	         "clientProvidedHostHeader": "ssm.sagov-westsouth-1r.amazonaws.com",
	         "tlsVersion": "TLSv1.2"
	      },
	      "userAgent": "stratus-red-team_ea782787-a65d-4fc4-9fca-1c97869a9a25",
	      "userIdentity": {
	         "accessKeyId": "AKIAW9X2Q2U25SK79UCX",
	         "accountId": "056392974792",
	         "arn": "arn:aws:iam::056392974792:user/christophe",
	         "principalId": "AIDA10CZIPPG73T21TDI",
	         "type": "IAMUser",
	         "userName": "christophe"
	      }
	   },
	   {
	      "awsRegion": "sagov-westsouth-1r",
	      "eventCategory": "Management",
	      "eventID": "bab0e5ba-5a43-467d-9460-dd801d9e9ad8",
	      "eventName": "GetCommandInvocation",
	      "eventSource": "ssm.amazonaws.com",
	      "eventTime": "2024-08-02T09:09:02Z",
	      "eventType": "AwsApiCall",
	      "eventVersion": "1.08",
	      "managementEvent": true,
	      "readOnly": true,
	      "recipientAccountId": "056392974792",
	      "requestID": "55198b26-f77b-4ef8-9259-bb347696f512",
	      "requestParameters": {
	         "commandId": "4e973221-443e-4a56-a0b4-1cb3c7923fc3",
	         "instanceId": "i-9D40CCFc0aE91CFa5"
	      },
	      "responseElements": null,
	      "sourceIPAddress": "253.252.51.07",
	      "tlsDetails": {
	         "cipherSuite": "ECDHE-RSA-AES128-GCM-SHA256",
	         "clientProvidedHostHeader": "ssm.sagov-westsouth-1r.amazonaws.com",
	         "tlsVersion": "TLSv1.2"
	      },
	      "userAgent": "stratus-red-team_ea782787-a65d-4fc4-9fca-1c97869a9a25",
	      "userIdentity": {
	         "accessKeyId": "AKIAW9X2Q2U25SK79UCX",
	         "accountId": "056392974792",
	         "arn": "arn:aws:iam::056392974792:user/christophe",
	         "principalId": "AIDA10CZIPPG73T21TDI",
	         "type": "IAMUser",
	         "userName": "christophe"
	      }
	   },
	   {
	      "awsRegion": "sagov-westsouth-1r",
	      "eventCategory": "Management",
	      "eventID": "b2c7717c-e542-422f-a78d-590536c174cb",
	      "eventName": "GetCommandInvocation",
	      "eventSource": "ssm.amazonaws.com",
	      "eventTime": "2024-08-02T09:09:01Z",
	      "eventType": "AwsApiCall",
	      "eventVersion": "1.08",
	      "managementEvent": true,
	      "readOnly": true,
	      "recipientAccountId": "056392974792",
	      "requestID": "3a1aa185-9cc4-4d58-933c-c2a6ad37c730",
	      "requestParameters": {
	         "commandId": "4e973221-443e-4a56-a0b4-1cb3c7923fc3",
	         "instanceId": "i-00456A8D163f546Ff"
	      },
	      "responseElements": null,
	      "sourceIPAddress": "253.252.51.07",
	      "tlsDetails": {
	         "cipherSuite": "ECDHE-RSA-AES128-GCM-SHA256",
	         "clientProvidedHostHeader": "ssm.sagov-westsouth-1r.amazonaws.com",
	         "tlsVersion": "TLSv1.2"
	      },
	      "userAgent": "stratus-red-team_ea782787-a65d-4fc4-9fca-1c97869a9a25",
	      "userIdentity": {
	         "accessKeyId": "AKIAW9X2Q2U25SK79UCX",
	         "accountId": "056392974792",
	         "arn": "arn:aws:iam::056392974792:user/christophe",
	         "principalId": "AIDA10CZIPPG73T21TDI",
	         "type": "IAMUser",
	         "userName": "christophe"
	      }
	   },
	   {
	      "awsRegion": "sagov-westsouth-1r",
	      "eventCategory": "Management",
	      "eventID": "e0b17230-9c13-482a-a0f0-d93c6bd4fb8e",
	      "eventName": "GetCommandInvocation",
	      "eventSource": "ssm.amazonaws.com",
	      "eventTime": "2024-08-02T09:09:01Z",
	      "eventType": "AwsApiCall",
	      "eventVersion": "1.08",
	      "managementEvent": true,
	      "readOnly": true,
	      "recipientAccountId": "056392974792",
	      "requestID": "dd526977-54b5-4951-bdb4-b9e542af402b",
	      "requestParameters": {
	         "commandId": "4e973221-443e-4a56-a0b4-1cb3c7923fc3",
	         "instanceId": "i-cfE23b1a7ceba6f86"
	      },
	      "responseElements": null,
	      "sourceIPAddress": "253.252.51.07",
	      "tlsDetails": {
	         "cipherSuite": "ECDHE-RSA-AES128-GCM-SHA256",
	         "clientProvidedHostHeader": "ssm.sagov-westsouth-1r.amazonaws.com",
	         "tlsVersion": "TLSv1.2"
	      },
	      "userAgent": "stratus-red-team_ea782787-a65d-4fc4-9fca-1c97869a9a25",
	      "userIdentity": {
	         "accessKeyId": "AKIAW9X2Q2U25SK79UCX",
	         "accountId": "056392974792",
	         "arn": "arn:aws:iam::056392974792:user/christophe",
	         "principalId": "AIDA10CZIPPG73T21TDI",
	         "type": "IAMUser",
	         "userName": "christophe"
	      }
	   },
	   {
	      "awsRegion": "sagov-westsouth-1r",
	      "eventCategory": "Management",
	      "eventID": "5288bfb8-e3fa-4c41-be02-6853521afe8b",
	      "eventName": "SendCommand",
	      "eventSource": "ssm.amazonaws.com",
	      "eventTime": "2024-08-02T09:08:56Z",
	      "eventType": "AwsApiCall",
	      "eventVersion": "1.08",
	      "managementEvent": true,
	      "readOnly": false,
	      "recipientAccountId": "056392974792",
	      "requestID": "1479b5e1-9751-4bf1-b548-cdd8108e85a6",
	      "requestParameters": {
	         "documentName": "AWS-RunShellScript",
	         "instanceIds": [
	            "i-00456A8D163f546Ff",
	            "i-cfE23b1a7ceba6f86",
	            "i-9D40CCFc0aE91CFa5"
	         ],
	         "interactive": false,
	         "parameters": "HIDDEN_DUE_TO_SECURITY_REASONS"
	      },
	      "responseElements": {
	         "command": {
	            "alarmConfiguration": {
	               "alarms": [],
	               "ignorePollAlarmFailure": false
	            },
	            "clientName": "",
	            "clientSourceId": "",
	            "cloudWatchOutputConfig": {
	               "cloudWatchLogGroupName": "",
	               "cloudWatchOutputEnabled": false
	            },
	            "commandId": "4e973221-443e-4a56-a0b4-1cb3c7923fc3",
	            "comment": "",
	            "completedCount": 0,
	            "deliveryTimedOutCount": 0,
	            "documentName": "AWS-RunShellScript",
	            "documentVersion": "$DEFAULT",
	            "errorCount": 0,
	            "expiresAfter": "Aug 2, 2024, 11:08:56 AM",
	            "hasCancelCommandSignature": false,
	            "hasSendCommandSignature": false,
	            "instanceIds": [
	               "i-00456A8D163f546Ff",
	               "i-cfE23b1a7ceba6f86",
	               "i-9D40CCFc0aE91CFa5"
	            ],
	            "interactive": false,
	            "maxConcurrency": "50",
	            "maxErrors": "0",
	            "notificationConfig": {
	               "notificationArn": "",
	               "notificationEvents": [],
	               "notificationType": ""
	            },
	            "outputS3BucketName": "",
	            "outputS3KeyPrefix": "",
	            "outputS3Region": "sagov-westsouth-1r",
	            "parameters": "HIDDEN_DUE_TO_SECURITY_REASONS",
	            "requestedDateTime": "Aug 2, 2024, 9:08:56 AM",
	            "serviceRole": "",
	            "status": "Pending",
	            "statusDetails": "Pending",
	            "targetCount": 3,
	            "targets": [],
	            "timeoutSeconds": 3600,
	            "triggeredAlarms": []
	         }
	      },
	      "sourceIPAddress": "253.252.51.07",
	      "tlsDetails": {
	         "cipherSuite": "ECDHE-RSA-AES128-GCM-SHA256",
	         "clientProvidedHostHeader": "ssm.sagov-westsouth-1r.amazonaws.com",
	         "tlsVersion": "TLSv1.2"
	      },
	      "userAgent": "stratus-red-team_ea782787-a65d-4fc4-9fca-1c97869a9a25",
	      "userIdentity": {
	         "accessKeyId": "AKIAW9X2Q2U25SK79UCX",
	         "accountId": "056392974792",
	         "arn": "arn:aws:iam::056392974792:user/christophe",
	         "principalId": "AIDA10CZIPPG73T21TDI",
	         "type": "IAMUser",
	         "userName": "christophe"
	      }
	   },
	   {
	      "awsRegion": "sagov-westsouth-1r",
	      "eventCategory": "Management",
	      "eventID": "8e1d1d98-6f88-4ce9-8e62-c1ec1a598408",
	      "eventName": "DescribeInstanceInformation",
	      "eventSource": "ssm.amazonaws.com",
	      "eventTime": "2024-08-02T09:08:56Z",
	      "eventType": "AwsApiCall",
	      "eventVersion": "1.08",
	      "managementEvent": true,
	      "readOnly": true,
	      "recipientAccountId": "056392974792",
	      "requestID": "5e34f5e1-11f1-481f-a435-c6124bd640d2",
	      "requestParameters": {
	         "filters": [
	            {
	               "key": "InstanceIds",
	               "values": [
	                  "i-00456A8D163f546Ff",
	                  "i-cfE23b1a7ceba6f86",
	                  "i-9D40CCFc0aE91CFa5"
	               ]
	            }
	         ]
	      },
	      "responseElements": null,
	      "sourceIPAddress": "253.252.51.07",
	      "tlsDetails": {
	         "cipherSuite": "ECDHE-RSA-AES128-GCM-SHA256",
	         "clientProvidedHostHeader": "ssm.sagov-westsouth-1r.amazonaws.com",
	         "tlsVersion": "TLSv1.2"
	      },
	      "userAgent": "stratus-red-team_ea782787-a65d-4fc4-9fca-1c97869a9a25",
	      "userIdentity": {
	         "accessKeyId": "AKIAW9X2Q2U25SK79UCX",
	         "accountId": "056392974792",
	         "arn": "arn:aws:iam::056392974792:user/christophe",
	         "principalId": "AIDA10CZIPPG73T21TDI",
	         "type": "IAMUser",
	         "userName": "christophe"
	      }
	   },
	   {
	      "awsRegion": "sagov-westsouth-1r",
	      "eventCategory": "Management",
	      "eventID": "e470e8f0-fbf0-42c1-a751-b271929bfa22",
	      "eventName": "GetCommandInvocation",
	      "eventSource": "ssm.amazonaws.com",
	      "eventTime": "2024-08-02T09:08:56Z",
	      "eventType": "AwsApiCall",
	      "eventVersion": "1.08",
	      "managementEvent": true,
	      "readOnly": true,
	      "recipientAccountId": "056392974792",
	      "requestID": "c6b8d64a-b975-4306-a8ac-17671377c2af",
	      "requestParameters": {
	         "commandId": "4e973221-443e-4a56-a0b4-1cb3c7923fc3",
	         "instanceId": "i-00456A8D163f546Ff"
	      },
	      "responseElements": null,
	      "sourceIPAddress": "253.252.51.07",
	      "tlsDetails": {
	         "cipherSuite": "ECDHE-RSA-AES128-GCM-SHA256",
	         "clientProvidedHostHeader": "ssm.sagov-westsouth-1r.amazonaws.com",
	         "tlsVersion": "TLSv1.2"
	      },
	      "userAgent": "stratus-red-team_ea782787-a65d-4fc4-9fca-1c97869a9a25",
	      "userIdentity": {
	         "accessKeyId": "AKIAW9X2Q2U25SK79UCX",
	         "accountId": "056392974792",
	         "arn": "arn:aws:iam::056392974792:user/christophe",
	         "principalId": "AIDA10CZIPPG73T21TDI",
	         "type": "IAMUser",
	         "userName": "christophe"
	      }
	   },
	   {
	      "awsRegion": "sagov-westsouth-1r",
	      "eventCategory": "Management",
	      "eventID": "ad342d3d-e850-41c3-b3a6-3e5cf0b382d3",
	      "eventName": "DescribeInstanceInformation",
	      "eventSource": "ssm.amazonaws.com",
	      "eventTime": "2024-08-02T09:08:55Z",
	      "eventType": "AwsApiCall",
	      "eventVersion": "1.08",
	      "managementEvent": true,
	      "readOnly": true,
	      "recipientAccountId": "056392974792",
	      "requestID": "6fd7d6fe-4452-462c-bf9c-c93daec119d6",
	      "requestParameters": {
	         "filters": [
	            {
	               "key": "InstanceIds",
	               "values": [
	                  "i-00456A8D163f546Ff",
	                  "i-cfE23b1a7ceba6f86",
	                  "i-9D40CCFc0aE91CFa5"
	               ]
	            }
	         ]
	      },
	      "responseElements": null,
	      "sourceIPAddress": "253.252.51.07",
	      "tlsDetails": {
	         "cipherSuite": "ECDHE-RSA-AES128-GCM-SHA256",
	         "clientProvidedHostHeader": "ssm.sagov-westsouth-1r.amazonaws.com",
	         "tlsVersion": "TLSv1.2"
	      },
	      "userAgent": "stratus-red-team_ea782787-a65d-4fc4-9fca-1c97869a9a25",
	      "userIdentity": {
	         "accessKeyId": "AKIAW9X2Q2U25SK79UCX",
	         "accountId": "056392974792",
	         "arn": "arn:aws:iam::056392974792:user/christophe",
	         "principalId": "AIDA10CZIPPG73T21TDI",
	         "type": "IAMUser",
	         "userName": "christophe"
	      }
	   },
	   {
	      "awsRegion": "sagov-westsouth-1r",
	      "eventCategory": "Management",
	      "eventID": "486ae737-1798-4c36-a90a-20d61f22d678",
	      "eventName": "DescribeInstanceInformation",
	      "eventSource": "ssm.amazonaws.com",
	      "eventTime": "2024-08-02T09:08:53Z",
	      "eventType": "AwsApiCall",
	      "eventVersion": "1.08",
	      "managementEvent": true,
	      "readOnly": true,
	      "recipientAccountId": "056392974792",
	      "requestID": "4dd32dc2-26bc-4d9a-a469-56c65a55f45e",
	      "requestParameters": {
	         "filters": [
	            {
	               "key": "InstanceIds",
	               "values": [
	                  "i-00456A8D163f546Ff",
	                  "i-cfE23b1a7ceba6f86",
	                  "i-9D40CCFc0aE91CFa5"
	               ]
	            }
	         ]
	      },
	      "responseElements": null,
	      "sourceIPAddress": "253.252.51.07",
	      "tlsDetails": {
	         "cipherSuite": "ECDHE-RSA-AES128-GCM-SHA256",
	         "clientProvidedHostHeader": "ssm.sagov-westsouth-1r.amazonaws.com",
	         "tlsVersion": "TLSv1.2"
	      },
	      "userAgent": "stratus-red-team_ea782787-a65d-4fc4-9fca-1c97869a9a25",
	      "userIdentity": {
	         "accessKeyId": "AKIAW9X2Q2U25SK79UCX",
	         "accountId": "056392974792",
	         "arn": "arn:aws:iam::056392974792:user/christophe",
	         "principalId": "AIDA10CZIPPG73T21TDI",
	         "type": "IAMUser",
	         "userName": "christophe"
	      }
	   },
	   {
	      "awsRegion": "sagov-westsouth-1r",
	      "eventCategory": "Management",
	      "eventID": "6643948a-9472-4f72-b1ff-8ddcfedca235",
	      "eventName": "DescribeInstanceInformation",
	      "eventSource": "ssm.amazonaws.com",
	      "eventTime": "2024-08-02T09:08:52Z",
	      "eventType": "AwsApiCall",
	      "eventVersion": "1.08",
	      "managementEvent": true,
	      "readOnly": true,
	      "recipientAccountId": "056392974792",
	      "requestID": "0605e0fd-df0a-493a-a915-832b50c17164",
	      "requestParameters": {
	         "filters": [
	            {
	               "key": "InstanceIds",
	               "values": [
	                  "i-00456A8D163f546Ff",
	                  "i-cfE23b1a7ceba6f86",
	                  "i-9D40CCFc0aE91CFa5"
	               ]
	            }
	         ]
	      },
	      "responseElements": null,
	      "sourceIPAddress": "253.252.51.07",
	      "tlsDetails": {
	         "cipherSuite": "ECDHE-RSA-AES128-GCM-SHA256",
	         "clientProvidedHostHeader": "ssm.sagov-westsouth-1r.amazonaws.com",
	         "tlsVersion": "TLSv1.2"
	      },
	      "userAgent": "stratus-red-team_ea782787-a65d-4fc4-9fca-1c97869a9a25",
	      "userIdentity": {
	         "accessKeyId": "AKIAW9X2Q2U25SK79UCX",
	         "accountId": "056392974792",
	         "arn": "arn:aws:iam::056392974792:user/christophe",
	         "principalId": "AIDA10CZIPPG73T21TDI",
	         "type": "IAMUser",
	         "userName": "christophe"
	      }
	   },
	   {
	      "awsRegion": "sagov-westsouth-1r",
	      "eventCategory": "Management",
	      "eventID": "cd49199d-ffdc-46bf-acae-e6c6d73e215a",
	      "eventName": "DescribeInstanceInformation",
	      "eventSource": "ssm.amazonaws.com",
	      "eventTime": "2024-08-02T09:08:51Z",
	      "eventType": "AwsApiCall",
	      "eventVersion": "1.08",
	      "managementEvent": true,
	      "readOnly": true,
	      "recipientAccountId": "056392974792",
	      "requestID": "65bc968b-731a-4dd5-93aa-3bfebcf16f85",
	      "requestParameters": {
	         "filters": [
	            {
	               "key": "InstanceIds",
	               "values": [
	                  "i-00456A8D163f546Ff",
	                  "i-cfE23b1a7ceba6f86",
	                  "i-9D40CCFc0aE91CFa5"
	               ]
	            }
	         ]
	      },
	      "responseElements": null,
	      "sourceIPAddress": "253.252.51.07",
	      "tlsDetails": {
	         "cipherSuite": "ECDHE-RSA-AES128-GCM-SHA256",
	         "clientProvidedHostHeader": "ssm.sagov-westsouth-1r.amazonaws.com",
	         "tlsVersion": "TLSv1.2"
	      },
	      "userAgent": "stratus-red-team_ea782787-a65d-4fc4-9fca-1c97869a9a25",
	      "userIdentity": {
	         "accessKeyId": "AKIAW9X2Q2U25SK79UCX",
	         "accountId": "056392974792",
	         "arn": "arn:aws:iam::056392974792:user/christophe",
	         "principalId": "AIDA10CZIPPG73T21TDI",
	         "type": "IAMUser",
	         "userName": "christophe"
	      }
	   },
	   {
	      "awsRegion": "sagov-westsouth-1r",
	      "eventCategory": "Management",
	      "eventID": "53407d54-9944-4317-a20f-d9a52c2a35ee",
	      "eventName": "DescribeInstanceInformation",
	      "eventSource": "ssm.amazonaws.com",
	      "eventTime": "2024-08-02T09:08:50Z",
	      "eventType": "AwsApiCall",
	      "eventVersion": "1.08",
	      "managementEvent": true,
	      "readOnly": true,
	      "recipientAccountId": "056392974792",
	      "requestID": "80ee2eb6-d794-4ac3-b2fb-6b9b40936d61",
	      "requestParameters": {
	         "filters": [
	            {
	               "key": "InstanceIds",
	               "values": [
	                  "i-00456A8D163f546Ff",
	                  "i-cfE23b1a7ceba6f86",
	                  "i-9D40CCFc0aE91CFa5"
	               ]
	            }
	         ]
	      },
	      "responseElements": null,
	      "sourceIPAddress": "253.252.51.07",
	      "tlsDetails": {
	         "cipherSuite": "ECDHE-RSA-AES128-GCM-SHA256",
	         "clientProvidedHostHeader": "ssm.sagov-westsouth-1r.amazonaws.com",
	         "tlsVersion": "TLSv1.2"
	      },
	      "userAgent": "stratus-red-team_ea782787-a65d-4fc4-9fca-1c97869a9a25",
	      "userIdentity": {
	         "accessKeyId": "AKIAW9X2Q2U25SK79UCX",
	         "accountId": "056392974792",
	         "arn": "arn:aws:iam::056392974792:user/christophe",
	         "principalId": "AIDA10CZIPPG73T21TDI",
	         "type": "IAMUser",
	         "userName": "christophe"
	      }
	   },
	   {
	      "awsRegion": "sagov-westsouth-1r",
	      "eventCategory": "Management",
	      "eventID": "6f1a2b4e-89a5-43f0-8ef4-6f3ecd9e04dc",
	      "eventName": "DescribeInstanceInformation",
	      "eventSource": "ssm.amazonaws.com",
	      "eventTime": "2024-08-02T09:08:49Z",
	      "eventType": "AwsApiCall",
	      "eventVersion": "1.08",
	      "managementEvent": true,
	      "readOnly": true,
	      "recipientAccountId": "056392974792",
	      "requestID": "5a765f60-eddc-4efe-bb7f-57b018f5c76a",
	      "requestParameters": {
	         "filters": [
	            {
	               "key": "InstanceIds",
	               "values": [
	                  "i-00456A8D163f546Ff",
	                  "i-cfE23b1a7ceba6f86",
	                  "i-9D40CCFc0aE91CFa5"
	               ]
	            }
	         ]
	      },
	      "responseElements": null,
	      "sourceIPAddress": "253.252.51.07",
	      "tlsDetails": {
	         "cipherSuite": "ECDHE-RSA-AES128-GCM-SHA256",
	         "clientProvidedHostHeader": "ssm.sagov-westsouth-1r.amazonaws.com",
	         "tlsVersion": "TLSv1.2"
	      },
	      "userAgent": "stratus-red-team_ea782787-a65d-4fc4-9fca-1c97869a9a25",
	      "userIdentity": {
	         "accessKeyId": "AKIAW9X2Q2U25SK79UCX",
	         "accountId": "056392974792",
	         "arn": "arn:aws:iam::056392974792:user/christophe",
	         "principalId": "AIDA10CZIPPG73T21TDI",
	         "type": "IAMUser",
	         "userName": "christophe"
	      }
	   },
	   {
	      "awsRegion": "sagov-westsouth-1r",
	      "eventCategory": "Management",
	      "eventID": "fdcf7d26-3ffb-4e35-8534-933b6ced55b5",
	      "eventName": "DescribeInstanceInformation",
	      "eventSource": "ssm.amazonaws.com",
	      "eventTime": "2024-08-02T09:08:48Z",
	      "eventType": "AwsApiCall",
	      "eventVersion": "1.08",
	      "managementEvent": true,
	      "readOnly": true,
	      "recipientAccountId": "056392974792",
	      "requestID": "42651f04-5238-4f63-889b-bee7734d29e0",
	      "requestParameters": {
	         "filters": [
	            {
	               "key": "InstanceIds",
	               "values": [
	                  "i-00456A8D163f546Ff",
	                  "i-cfE23b1a7ceba6f86",
	                  "i-9D40CCFc0aE91CFa5"
	               ]
	            }
	         ]
	      },
	      "responseElements": null,
	      "sourceIPAddress": "253.252.51.07",
	      "tlsDetails": {
	         "cipherSuite": "ECDHE-RSA-AES128-GCM-SHA256",
	         "clientProvidedHostHeader": "ssm.sagov-westsouth-1r.amazonaws.com",
	         "tlsVersion": "TLSv1.2"
	      },
	      "userAgent": "stratus-red-team_ea782787-a65d-4fc4-9fca-1c97869a9a25",
	      "userIdentity": {
	         "accessKeyId": "AKIAW9X2Q2U25SK79UCX",
	         "accountId": "056392974792",
	         "arn": "arn:aws:iam::056392974792:user/christophe",
	         "principalId": "AIDA10CZIPPG73T21TDI",
	         "type": "IAMUser",
	         "userName": "christophe"
	      }
	   },
	   {
	      "awsRegion": "sagov-westsouth-1r",
	      "eventCategory": "Management",
	      "eventID": "1a5374a3-1223-46dc-b3c4-a0336179f22b",
	      "eventName": "DescribeInstanceInformation",
	      "eventSource": "ssm.amazonaws.com",
	      "eventTime": "2024-08-02T09:08:46Z",
	      "eventType": "AwsApiCall",
	      "eventVersion": "1.08",
	      "managementEvent": true,
	      "readOnly": true,
	      "recipientAccountId": "056392974792",
	      "requestID": "f12f2209-52ba-4064-8e48-45a70ed55437",
	      "requestParameters": {
	         "filters": [
	            {
	               "key": "InstanceIds",
	               "values": [
	                  "i-00456A8D163f546Ff",
	                  "i-cfE23b1a7ceba6f86",
	                  "i-9D40CCFc0aE91CFa5"
	               ]
	            }
	         ]
	      },
	      "responseElements": null,
	      "sourceIPAddress": "253.252.51.07",
	      "tlsDetails": {
	         "cipherSuite": "ECDHE-RSA-AES128-GCM-SHA256",
	         "clientProvidedHostHeader": "ssm.sagov-westsouth-1r.amazonaws.com",
	         "tlsVersion": "TLSv1.2"
	      },
	      "userAgent": "stratus-red-team_ea782787-a65d-4fc4-9fca-1c97869a9a25",
	      "userIdentity": {
	         "accessKeyId": "AKIAW9X2Q2U25SK79UCX",
	         "accountId": "056392974792",
	         "arn": "arn:aws:iam::056392974792:user/christophe",
	         "principalId": "AIDA10CZIPPG73T21TDI",
	         "type": "IAMUser",
	         "userName": "christophe"
	      }
	   },
	   {
	      "awsRegion": "sagov-westsouth-1r",
	      "eventCategory": "Management",
	      "eventID": "1fc0903a-bdd5-4a31-a15e-84efb05530dd",
	      "eventName": "DescribeInstanceInformation",
	      "eventSource": "ssm.amazonaws.com",
	      "eventTime": "2024-08-02T09:08:45Z",
	      "eventType": "AwsApiCall",
	      "eventVersion": "1.08",
	      "managementEvent": true,
	      "readOnly": true,
	      "recipientAccountId": "056392974792",
	      "requestID": "54a4713e-2480-4b3c-95de-ffa6f061f6db",
	      "requestParameters": {
	         "filters": [
	            {
	               "key": "InstanceIds",
	               "values": [
	                  "i-00456A8D163f546Ff",
	                  "i-cfE23b1a7ceba6f86",
	                  "i-9D40CCFc0aE91CFa5"
	               ]
	            }
	         ]
	      },
	      "responseElements": null,
	      "sourceIPAddress": "253.252.51.07",
	      "tlsDetails": {
	         "cipherSuite": "ECDHE-RSA-AES128-GCM-SHA256",
	         "clientProvidedHostHeader": "ssm.sagov-westsouth-1r.amazonaws.com",
	         "tlsVersion": "TLSv1.2"
	      },
	      "userAgent": "stratus-red-team_ea782787-a65d-4fc4-9fca-1c97869a9a25",
	      "userIdentity": {
	         "accessKeyId": "AKIAW9X2Q2U25SK79UCX",
	         "accountId": "056392974792",
	         "arn": "arn:aws:iam::056392974792:user/christophe",
	         "principalId": "AIDA10CZIPPG73T21TDI",
	         "type": "IAMUser",
	         "userName": "christophe"
	      }
	   },
	   {
	      "awsRegion": "sagov-westsouth-1r",
	      "eventCategory": "Management",
	      "eventID": "b43fdb25-5caf-4203-b2f4-5fd4d40344b0",
	      "eventName": "DescribeInstanceInformation",
	      "eventSource": "ssm.amazonaws.com",
	      "eventTime": "2024-08-02T09:08:44Z",
	      "eventType": "AwsApiCall",
	      "eventVersion": "1.08",
	      "managementEvent": true,
	      "readOnly": true,
	      "recipientAccountId": "056392974792",
	      "requestID": "c2342054-aa38-41f4-b1b9-702828726730",
	      "requestParameters": {
	         "filters": [
	            {
	               "key": "InstanceIds",
	               "values": [
	                  "i-00456A8D163f546Ff",
	                  "i-cfE23b1a7ceba6f86",
	                  "i-9D40CCFc0aE91CFa5"
	               ]
	            }
	         ]
	      },
	      "responseElements": null,
	      "sourceIPAddress": "253.252.51.07",
	      "tlsDetails": {
	         "cipherSuite": "ECDHE-RSA-AES128-GCM-SHA256",
	         "clientProvidedHostHeader": "ssm.sagov-westsouth-1r.amazonaws.com",
	         "tlsVersion": "TLSv1.2"
	      },
	      "userAgent": "stratus-red-team_ea782787-a65d-4fc4-9fca-1c97869a9a25",
	      "userIdentity": {
	         "accessKeyId": "AKIAW9X2Q2U25SK79UCX",
	         "accountId": "056392974792",
	         "arn": "arn:aws:iam::056392974792:user/christophe",
	         "principalId": "AIDA10CZIPPG73T21TDI",
	         "type": "IAMUser",
	         "userName": "christophe"
	      }
	   },
	   {
	      "awsRegion": "sagov-westsouth-1r",
	      "eventCategory": "Management",
	      "eventID": "97a253c0-5e84-4d78-8412-a420695ba4dc",
	      "eventName": "DescribeInstanceInformation",
	      "eventSource": "ssm.amazonaws.com",
	      "eventTime": "2024-08-02T09:08:43Z",
	      "eventType": "AwsApiCall",
	      "eventVersion": "1.08",
	      "managementEvent": true,
	      "readOnly": true,
	      "recipientAccountId": "056392974792",
	      "requestID": "398704b7-2c17-4cb2-8efb-f27ef8f775fe",
	      "requestParameters": {
	         "filters": [
	            {
	               "key": "InstanceIds",
	               "values": [
	                  "i-00456A8D163f546Ff",
	                  "i-cfE23b1a7ceba6f86",
	                  "i-9D40CCFc0aE91CFa5"
	               ]
	            }
	         ]
	      },
	      "responseElements": null,
	      "sourceIPAddress": "253.252.51.07",
	      "tlsDetails": {
	         "cipherSuite": "ECDHE-RSA-AES128-GCM-SHA256",
	         "clientProvidedHostHeader": "ssm.sagov-westsouth-1r.amazonaws.com",
	         "tlsVersion": "TLSv1.2"
	      },
	      "userAgent": "stratus-red-team_ea782787-a65d-4fc4-9fca-1c97869a9a25",
	      "userIdentity": {
	         "accessKeyId": "AKIAW9X2Q2U25SK79UCX",
	         "accountId": "056392974792",
	         "arn": "arn:aws:iam::056392974792:user/christophe",
	         "principalId": "AIDA10CZIPPG73T21TDI",
	         "type": "IAMUser",
	         "userName": "christophe"
	      }
	   },
	   {
	      "awsRegion": "sagov-westsouth-1r",
	      "eventCategory": "Management",
	      "eventID": "e4be349d-0420-4ee9-b8da-7f8b76c4d883",
	      "eventName": "DescribeInstanceInformation",
	      "eventSource": "ssm.amazonaws.com",
	      "eventTime": "2024-08-02T09:08:42Z",
	      "eventType": "AwsApiCall",
	      "eventVersion": "1.08",
	      "managementEvent": true,
	      "readOnly": true,
	      "recipientAccountId": "056392974792",
	      "requestID": "5db544de-5064-4bf2-ba19-ea2a882281bc",
	      "requestParameters": {
	         "filters": [
	            {
	               "key": "InstanceIds",
	               "values": [
	                  "i-00456A8D163f546Ff",
	                  "i-cfE23b1a7ceba6f86",
	                  "i-9D40CCFc0aE91CFa5"
	               ]
	            }
	         ]
	      },
	      "responseElements": null,
	      "sourceIPAddress": "253.252.51.07",
	      "tlsDetails": {
	         "cipherSuite": "ECDHE-RSA-AES128-GCM-SHA256",
	         "clientProvidedHostHeader": "ssm.sagov-westsouth-1r.amazonaws.com",
	         "tlsVersion": "TLSv1.2"
	      },
	      "userAgent": "stratus-red-team_ea782787-a65d-4fc4-9fca-1c97869a9a25",
	      "userIdentity": {
	         "accessKeyId": "AKIAW9X2Q2U25SK79UCX",
	         "accountId": "056392974792",
	         "arn": "arn:aws:iam::056392974792:user/christophe",
	         "principalId": "AIDA10CZIPPG73T21TDI",
	         "type": "IAMUser",
	         "userName": "christophe"
	      }
	   },
	   {
	      "awsRegion": "sagov-westsouth-1r",
	      "eventCategory": "Management",
	      "eventID": "55b6e5a7-e4e8-4b81-b822-75905525c193",
	      "eventName": "DescribeInstanceInformation",
	      "eventSource": "ssm.amazonaws.com",
	      "eventTime": "2024-08-02T09:08:41Z",
	      "eventType": "AwsApiCall",
	      "eventVersion": "1.08",
	      "managementEvent": true,
	      "readOnly": true,
	      "recipientAccountId": "056392974792",
	      "requestID": "2654285f-1d76-4224-9224-4a3968f16a3f",
	      "requestParameters": {
	         "filters": [
	            {
	               "key": "InstanceIds",
	               "values": [
	                  "i-00456A8D163f546Ff",
	                  "i-cfE23b1a7ceba6f86",
	                  "i-9D40CCFc0aE91CFa5"
	               ]
	            }
	         ]
	      },
	      "responseElements": null,
	      "sourceIPAddress": "253.252.51.07",
	      "tlsDetails": {
	         "cipherSuite": "ECDHE-RSA-AES128-GCM-SHA256",
	         "clientProvidedHostHeader": "ssm.sagov-westsouth-1r.amazonaws.com",
	         "tlsVersion": "TLSv1.2"
	      },
	      "userAgent": "stratus-red-team_ea782787-a65d-4fc4-9fca-1c97869a9a25",
	      "userIdentity": {
	         "accessKeyId": "AKIAW9X2Q2U25SK79UCX",
	         "accountId": "056392974792",
	         "arn": "arn:aws:iam::056392974792:user/christophe",
	         "principalId": "AIDA10CZIPPG73T21TDI",
	         "type": "IAMUser",
	         "userName": "christophe"
	      }
	   },
	   {
	      "awsRegion": "sagov-westsouth-1r",
	      "eventCategory": "Management",
	      "eventID": "c0679959-5bf1-4aaf-9f78-f436c35da4b2",
	      "eventName": "DescribeInstanceInformation",
	      "eventSource": "ssm.amazonaws.com",
	      "eventTime": "2024-08-02T09:08:39Z",
	      "eventType": "AwsApiCall",
	      "eventVersion": "1.08",
	      "managementEvent": true,
	      "readOnly": true,
	      "recipientAccountId": "056392974792",
	      "requestID": "1545c090-8ecf-4cae-9db0-a2da1e103f23",
	      "requestParameters": {
	         "filters": [
	            {
	               "key": "InstanceIds",
	               "values": [
	                  "i-00456A8D163f546Ff",
	                  "i-cfE23b1a7ceba6f86",
	                  "i-9D40CCFc0aE91CFa5"
	               ]
	            }
	         ]
	      },
	      "responseElements": null,
	      "sourceIPAddress": "253.252.51.07",
	      "tlsDetails": {
	         "cipherSuite": "ECDHE-RSA-AES128-GCM-SHA256",
	         "clientProvidedHostHeader": "ssm.sagov-westsouth-1r.amazonaws.com",
	         "tlsVersion": "TLSv1.2"
	      },
	      "userAgent": "stratus-red-team_ea782787-a65d-4fc4-9fca-1c97869a9a25",
	      "userIdentity": {
	         "accessKeyId": "AKIAW9X2Q2U25SK79UCX",
	         "accountId": "056392974792",
	         "arn": "arn:aws:iam::056392974792:user/christophe",
	         "principalId": "AIDA10CZIPPG73T21TDI",
	         "type": "IAMUser",
	         "userName": "christophe"
	      }
	   },
	   {
	      "awsRegion": "sagov-westsouth-1r",
	      "eventCategory": "Management",
	      "eventID": "bf330a73-3600-4a88-a3c9-837c82fd6431",
	      "eventName": "DescribeInstanceInformation",
	      "eventSource": "ssm.amazonaws.com",
	      "eventTime": "2024-08-02T09:08:38Z",
	      "eventType": "AwsApiCall",
	      "eventVersion": "1.08",
	      "managementEvent": true,
	      "readOnly": true,
	      "recipientAccountId": "056392974792",
	      "requestID": "409166a6-71c7-4a1c-b1dd-7972ec637a0c",
	      "requestParameters": {
	         "filters": [
	            {
	               "key": "InstanceIds",
	               "values": [
	                  "i-00456A8D163f546Ff",
	                  "i-cfE23b1a7ceba6f86",
	                  "i-9D40CCFc0aE91CFa5"
	               ]
	            }
	         ]
	      },
	      "responseElements": null,
	      "sourceIPAddress": "253.252.51.07",
	      "tlsDetails": {
	         "cipherSuite": "ECDHE-RSA-AES128-GCM-SHA256",
	         "clientProvidedHostHeader": "ssm.sagov-westsouth-1r.amazonaws.com",
	         "tlsVersion": "TLSv1.2"
	      },
	      "userAgent": "stratus-red-team_ea782787-a65d-4fc4-9fca-1c97869a9a25",
	      "userIdentity": {
	         "accessKeyId": "AKIAW9X2Q2U25SK79UCX",
	         "accountId": "056392974792",
	         "arn": "arn:aws:iam::056392974792:user/christophe",
	         "principalId": "AIDA10CZIPPG73T21TDI",
	         "type": "IAMUser",
	         "userName": "christophe"
	      }
	   },
	   {
	      "awsRegion": "sagov-westsouth-1r",
	      "eventCategory": "Management",
	      "eventID": "d303c923-1ad3-4333-a78c-5ba0d713df14",
	      "eventName": "DescribeInstanceInformation",
	      "eventSource": "ssm.amazonaws.com",
	      "eventTime": "2024-08-02T09:08:37Z",
	      "eventType": "AwsApiCall",
	      "eventVersion": "1.08",
	      "managementEvent": true,
	      "readOnly": true,
	      "recipientAccountId": "056392974792",
	      "requestID": "29eb2c6a-3d0a-4b1c-b643-ad80f5faee5f",
	      "requestParameters": {
	         "filters": [
	            {
	               "key": "InstanceIds",
	               "values": [
	                  "i-00456A8D163f546Ff",
	                  "i-cfE23b1a7ceba6f86",
	                  "i-9D40CCFc0aE91CFa5"
	               ]
	            }
	         ]
	      },
	      "responseElements": null,
	      "sourceIPAddress": "253.252.51.07",
	      "tlsDetails": {
	         "cipherSuite": "ECDHE-RSA-AES128-GCM-SHA256",
	         "clientProvidedHostHeader": "ssm.sagov-westsouth-1r.amazonaws.com",
	         "tlsVersion": "TLSv1.2"
	      },
	      "userAgent": "stratus-red-team_ea782787-a65d-4fc4-9fca-1c97869a9a25",
	      "userIdentity": {
	         "accessKeyId": "AKIAW9X2Q2U25SK79UCX",
	         "accountId": "056392974792",
	         "arn": "arn:aws:iam::056392974792:user/christophe",
	         "principalId": "AIDA10CZIPPG73T21TDI",
	         "type": "IAMUser",
	         "userName": "christophe"
	      }
	   },
	   {
	      "awsRegion": "sagov-westsouth-1r",
	      "eventCategory": "Management",
	      "eventID": "7cf67dfd-fedc-4494-acbe-3fab7e1808a1",
	      "eventName": "DescribeInstanceInformation",
	      "eventSource": "ssm.amazonaws.com",
	      "eventTime": "2024-08-02T09:08:36Z",
	      "eventType": "AwsApiCall",
	      "eventVersion": "1.08",
	      "managementEvent": true,
	      "readOnly": true,
	      "recipientAccountId": "056392974792",
	      "requestID": "9525e5ee-669c-40a2-a8d2-33cebb0ee895",
	      "requestParameters": {
	         "filters": [
	            {
	               "key": "InstanceIds",
	               "values": [
	                  "i-00456A8D163f546Ff",
	                  "i-cfE23b1a7ceba6f86",
	                  "i-9D40CCFc0aE91CFa5"
	               ]
	            }
	         ]
	      },
	      "responseElements": null,
	      "sourceIPAddress": "253.252.51.07",
	      "tlsDetails": {
	         "cipherSuite": "ECDHE-RSA-AES128-GCM-SHA256",
	         "clientProvidedHostHeader": "ssm.sagov-westsouth-1r.amazonaws.com",
	         "tlsVersion": "TLSv1.2"
	      },
	      "userAgent": "stratus-red-team_ea782787-a65d-4fc4-9fca-1c97869a9a25",
	      "userIdentity": {
	         "accessKeyId": "AKIAW9X2Q2U25SK79UCX",
	         "accountId": "056392974792",
	         "arn": "arn:aws:iam::056392974792:user/christophe",
	         "principalId": "AIDA10CZIPPG73T21TDI",
	         "type": "IAMUser",
	         "userName": "christophe"
	      }
	   },
	   {
	      "awsRegion": "sagov-westsouth-1r",
	      "eventCategory": "Management",
	      "eventID": "e666a3d4-db2f-4ac2-b0ba-63531a949154",
	      "eventName": "DescribeInstanceInformation",
	      "eventSource": "ssm.amazonaws.com",
	      "eventTime": "2024-08-02T09:08:35Z",
	      "eventType": "AwsApiCall",
	      "eventVersion": "1.08",
	      "managementEvent": true,
	      "readOnly": true,
	      "recipientAccountId": "056392974792",
	      "requestID": "34ea6034-0028-46cd-94f5-54ffb4c5ba02",
	      "requestParameters": {
	         "filters": [
	            {
	               "key": "InstanceIds",
	               "values": [
	                  "i-00456A8D163f546Ff",
	                  "i-cfE23b1a7ceba6f86",
	                  "i-9D40CCFc0aE91CFa5"
	               ]
	            }
	         ]
	      },
	      "responseElements": null,
	      "sourceIPAddress": "253.252.51.07",
	      "tlsDetails": {
	         "cipherSuite": "ECDHE-RSA-AES128-GCM-SHA256",
	         "clientProvidedHostHeader": "ssm.sagov-westsouth-1r.amazonaws.com",
	         "tlsVersion": "TLSv1.2"
	      },
	      "userAgent": "stratus-red-team_ea782787-a65d-4fc4-9fca-1c97869a9a25",
	      "userIdentity": {
	         "accessKeyId": "AKIAW9X2Q2U25SK79UCX",
	         "accountId": "056392974792",
	         "arn": "arn:aws:iam::056392974792:user/christophe",
	         "principalId": "AIDA10CZIPPG73T21TDI",
	         "type": "IAMUser",
	         "userName": "christophe"
	      }
	   },
	   {
	      "awsRegion": "sagov-westsouth-1r",
	      "eventCategory": "Management",
	      "eventID": "ff0452d7-bef3-47ba-b641-e4b10f50f3c4",
	      "eventName": "DescribeInstanceInformation",
	      "eventSource": "ssm.amazonaws.com",
	      "eventTime": "2024-08-02T09:08:34Z",
	      "eventType": "AwsApiCall",
	      "eventVersion": "1.08",
	      "managementEvent": true,
	      "readOnly": true,
	      "recipientAccountId": "056392974792",
	      "requestID": "135ea4ff-0e59-4771-b541-326b904dfd70",
	      "requestParameters": {
	         "filters": [
	            {
	               "key": "InstanceIds",
	               "values": [
	                  "i-00456A8D163f546Ff",
	                  "i-cfE23b1a7ceba6f86",
	                  "i-9D40CCFc0aE91CFa5"
	               ]
	            }
	         ]
	      },
	      "responseElements": null,
	      "sourceIPAddress": "253.252.51.07",
	      "tlsDetails": {
	         "cipherSuite": "ECDHE-RSA-AES128-GCM-SHA256",
	         "clientProvidedHostHeader": "ssm.sagov-westsouth-1r.amazonaws.com",
	         "tlsVersion": "TLSv1.2"
	      },
	      "userAgent": "stratus-red-team_ea782787-a65d-4fc4-9fca-1c97869a9a25",
	      "userIdentity": {
	         "accessKeyId": "AKIAW9X2Q2U25SK79UCX",
	         "accountId": "056392974792",
	         "arn": "arn:aws:iam::056392974792:user/christophe",
	         "principalId": "AIDA10CZIPPG73T21TDI",
	         "type": "IAMUser",
	         "userName": "christophe"
	      }
	   },
	   {
	      "awsRegion": "sagov-westsouth-1r",
	      "eventCategory": "Management",
	      "eventID": "2253ede9-2382-41fa-8302-b25ecf0f11ac",
	      "eventName": "DescribeInstanceInformation",
	      "eventSource": "ssm.amazonaws.com",
	      "eventTime": "2024-08-02T09:08:33Z",
	      "eventType": "AwsApiCall",
	      "eventVersion": "1.08",
	      "managementEvent": true,
	      "readOnly": true,
	      "recipientAccountId": "056392974792",
	      "requestID": "0c664d14-0f8b-44da-896d-80b7dae05a2c",
	      "requestParameters": {
	         "filters": [
	            {
	               "key": "InstanceIds",
	               "values": [
	                  "i-00456A8D163f546Ff",
	                  "i-cfE23b1a7ceba6f86",
	                  "i-9D40CCFc0aE91CFa5"
	               ]
	            }
	         ]
	      },
	      "responseElements": null,
	      "sourceIPAddress": "253.252.51.07",
	      "tlsDetails": {
	         "cipherSuite": "ECDHE-RSA-AES128-GCM-SHA256",
	         "clientProvidedHostHeader": "ssm.sagov-westsouth-1r.amazonaws.com",
	         "tlsVersion": "TLSv1.2"
	      },
	      "userAgent": "stratus-red-team_ea782787-a65d-4fc4-9fca-1c97869a9a25",
	      "userIdentity": {
	         "accessKeyId": "AKIAW9X2Q2U25SK79UCX",
	         "accountId": "056392974792",
	         "arn": "arn:aws:iam::056392974792:user/christophe",
	         "principalId": "AIDA10CZIPPG73T21TDI",
	         "type": "IAMUser",
	         "userName": "christophe"
	      }
	   },
	   {
	      "awsRegion": "sagov-westsouth-1r",
	      "eventCategory": "Management",
	      "eventID": "9b6c78ee-98ba-4ddd-9dae-aa4d3a57e89c",
	      "eventName": "DescribeInstanceInformation",
	      "eventSource": "ssm.amazonaws.com",
	      "eventTime": "2024-08-02T09:08:31Z",
	      "eventType": "AwsApiCall",
	      "eventVersion": "1.08",
	      "managementEvent": true,
	      "readOnly": true,
	      "recipientAccountId": "056392974792",
	      "requestID": "f55872e1-6dad-42be-a18d-c7bd64ef9f6d",
	      "requestParameters": {
	         "filters": [
	            {
	               "key": "InstanceIds",
	               "values": [
	                  "i-00456A8D163f546Ff",
	                  "i-cfE23b1a7ceba6f86",
	                  "i-9D40CCFc0aE91CFa5"
	               ]
	            }
	         ]
	      },
	      "responseElements": null,
	      "sourceIPAddress": "253.252.51.07",
	      "tlsDetails": {
	         "cipherSuite": "ECDHE-RSA-AES128-GCM-SHA256",
	         "clientProvidedHostHeader": "ssm.sagov-westsouth-1r.amazonaws.com",
	         "tlsVersion": "TLSv1.2"
	      },
	      "userAgent": "stratus-red-team_ea782787-a65d-4fc4-9fca-1c97869a9a25",
	      "userIdentity": {
	         "accessKeyId": "AKIAW9X2Q2U25SK79UCX",
	         "accountId": "056392974792",
	         "arn": "arn:aws:iam::056392974792:user/christophe",
	         "principalId": "AIDA10CZIPPG73T21TDI",
	         "type": "IAMUser",
	         "userName": "christophe"
	      }
	   },
	   {
	      "awsRegion": "sagov-westsouth-1r",
	      "eventCategory": "Management",
	      "eventID": "1ac28c35-ee6f-41a4-97bd-ae8e44363660",
	      "eventName": "DescribeInstanceInformation",
	      "eventSource": "ssm.amazonaws.com",
	      "eventTime": "2024-08-02T09:08:30Z",
	      "eventType": "AwsApiCall",
	      "eventVersion": "1.08",
	      "managementEvent": true,
	      "readOnly": true,
	      "recipientAccountId": "056392974792",
	      "requestID": "c274e01e-2045-4415-bd71-c8744107618e",
	      "requestParameters": {
	         "filters": [
	            {
	               "key": "InstanceIds",
	               "values": [
	                  "i-00456A8D163f546Ff",
	                  "i-cfE23b1a7ceba6f86",
	                  "i-9D40CCFc0aE91CFa5"
	               ]
	            }
	         ]
	      },
	      "responseElements": null,
	      "sourceIPAddress": "253.252.51.07",
	      "tlsDetails": {
	         "cipherSuite": "ECDHE-RSA-AES128-GCM-SHA256",
	         "clientProvidedHostHeader": "ssm.sagov-westsouth-1r.amazonaws.com",
	         "tlsVersion": "TLSv1.2"
	      },
	      "userAgent": "stratus-red-team_ea782787-a65d-4fc4-9fca-1c97869a9a25",
	      "userIdentity": {
	         "accessKeyId": "AKIAW9X2Q2U25SK79UCX",
	         "accountId": "056392974792",
	         "arn": "arn:aws:iam::056392974792:user/christophe",
	         "principalId": "AIDA10CZIPPG73T21TDI",
	         "type": "IAMUser",
	         "userName": "christophe"
	      }
	   },
	   {
	      "awsRegion": "sagov-westsouth-1r",
	      "eventCategory": "Management",
	      "eventID": "d3471df2-fc63-479b-9920-4ac3c9c32357",
	      "eventName": "DescribeInstanceInformation",
	      "eventSource": "ssm.amazonaws.com",
	      "eventTime": "2024-08-02T09:08:29Z",
	      "eventType": "AwsApiCall",
	      "eventVersion": "1.08",
	      "managementEvent": true,
	      "readOnly": true,
	      "recipientAccountId": "056392974792",
	      "requestID": "00d4a58a-00a8-4116-b391-beaa8aa1c0db",
	      "requestParameters": {
	         "filters": [
	            {
	               "key": "InstanceIds",
	               "values": [
	                  "i-00456A8D163f546Ff",
	                  "i-cfE23b1a7ceba6f86",
	                  "i-9D40CCFc0aE91CFa5"
	               ]
	            }
	         ]
	      },
	      "responseElements": null,
	      "sourceIPAddress": "253.252.51.07",
	      "tlsDetails": {
	         "cipherSuite": "ECDHE-RSA-AES128-GCM-SHA256",
	         "clientProvidedHostHeader": "ssm.sagov-westsouth-1r.amazonaws.com",
	         "tlsVersion": "TLSv1.2"
	      },
	      "userAgent": "stratus-red-team_ea782787-a65d-4fc4-9fca-1c97869a9a25",
	      "userIdentity": {
	         "accessKeyId": "AKIAW9X2Q2U25SK79UCX",
	         "accountId": "056392974792",
	         "arn": "arn:aws:iam::056392974792:user/christophe",
	         "principalId": "AIDA10CZIPPG73T21TDI",
	         "type": "IAMUser",
	         "userName": "christophe"
	      }
	   },
	   {
	      "awsRegion": "sagov-westsouth-1r",
	      "eventCategory": "Management",
	      "eventID": "0745f3f1-b181-4395-a2dc-243becae570e",
	      "eventName": "DescribeInstanceInformation",
	      "eventSource": "ssm.amazonaws.com",
	      "eventTime": "2024-08-02T09:08:28Z",
	      "eventType": "AwsApiCall",
	      "eventVersion": "1.08",
	      "managementEvent": true,
	      "readOnly": true,
	      "recipientAccountId": "056392974792",
	      "requestID": "4b2f5fd6-3620-4aa7-bf3e-7da9d27bec85",
	      "requestParameters": {
	         "filters": [
	            {
	               "key": "InstanceIds",
	               "values": [
	                  "i-00456A8D163f546Ff",
	                  "i-cfE23b1a7ceba6f86",
	                  "i-9D40CCFc0aE91CFa5"
	               ]
	            }
	         ]
	      },
	      "responseElements": null,
	      "sourceIPAddress": "253.252.51.07",
	      "tlsDetails": {
	         "cipherSuite": "ECDHE-RSA-AES128-GCM-SHA256",
	         "clientProvidedHostHeader": "ssm.sagov-westsouth-1r.amazonaws.com",
	         "tlsVersion": "TLSv1.2"
	      },
	      "userAgent": "stratus-red-team_ea782787-a65d-4fc4-9fca-1c97869a9a25",
	      "userIdentity": {
	         "accessKeyId": "AKIAW9X2Q2U25SK79UCX",
	         "accountId": "056392974792",
	         "arn": "arn:aws:iam::056392974792:user/christophe",
	         "principalId": "AIDA10CZIPPG73T21TDI",
	         "type": "IAMUser",
	         "userName": "christophe"
	      }
	   },
	   {
	      "awsRegion": "sagov-westsouth-1r",
	      "eventCategory": "Management",
	      "eventID": "804c4178-75cd-4d83-b04f-960f47961a75",
	      "eventName": "DescribeInstanceInformation",
	      "eventSource": "ssm.amazonaws.com",
	      "eventTime": "2024-08-02T09:08:24Z",
	      "eventType": "AwsApiCall",
	      "eventVersion": "1.08",
	      "managementEvent": true,
	      "readOnly": true,
	      "recipientAccountId": "056392974792",
	      "requestID": "bec61003-0f60-45c0-9256-116efb6d15aa",
	      "requestParameters": {
	         "filters": [
	            {
	               "key": "InstanceIds",
	               "values": [
	                  "i-00456A8D163f546Ff",
	                  "i-cfE23b1a7ceba6f86",
	                  "i-9D40CCFc0aE91CFa5"
	               ]
	            }
	         ]
	      },
	      "responseElements": null,
	      "sourceIPAddress": "253.252.51.07",
	      "tlsDetails": {
	         "cipherSuite": "ECDHE-RSA-AES128-GCM-SHA256",
	         "clientProvidedHostHeader": "ssm.sagov-westsouth-1r.amazonaws.com",
	         "tlsVersion": "TLSv1.2"
	      },
	      "userAgent": "stratus-red-team_ea782787-a65d-4fc4-9fca-1c97869a9a25",
	      "userIdentity": {
	         "accessKeyId": "AKIAW9X2Q2U25SK79UCX",
	         "accountId": "056392974792",
	         "arn": "arn:aws:iam::056392974792:user/christophe",
	         "principalId": "AIDA10CZIPPG73T21TDI",
	         "type": "IAMUser",
	         "userName": "christophe"
	      }
	   },
	   {
	      "awsRegion": "sagov-westsouth-1r",
	      "eventCategory": "Management",
	      "eventID": "73518501-d83c-4d7e-8dbd-2154928d76f7",
	      "eventName": "DescribeInstanceInformation",
	      "eventSource": "ssm.amazonaws.com",
	      "eventTime": "2024-08-02T09:08:22Z",
	      "eventType": "AwsApiCall",
	      "eventVersion": "1.08",
	      "managementEvent": true,
	      "readOnly": true,
	      "recipientAccountId": "056392974792",
	      "requestID": "4c950f64-59ff-4fce-9a69-32ef10f96872",
	      "requestParameters": {
	         "filters": [
	            {
	               "key": "InstanceIds",
	               "values": [
	                  "i-00456A8D163f546Ff",
	                  "i-cfE23b1a7ceba6f86",
	                  "i-9D40CCFc0aE91CFa5"
	               ]
	            }
	         ]
	      },
	      "responseElements": null,
	      "sourceIPAddress": "253.252.51.07",
	      "tlsDetails": {
	         "cipherSuite": "ECDHE-RSA-AES128-GCM-SHA256",
	         "clientProvidedHostHeader": "ssm.sagov-westsouth-1r.amazonaws.com",
	         "tlsVersion": "TLSv1.2"
	      },
	      "userAgent": "stratus-red-team_ea782787-a65d-4fc4-9fca-1c97869a9a25",
	      "userIdentity": {
	         "accessKeyId": "AKIAW9X2Q2U25SK79UCX",
	         "accountId": "056392974792",
	         "arn": "arn:aws:iam::056392974792:user/christophe",
	         "principalId": "AIDA10CZIPPG73T21TDI",
	         "type": "IAMUser",
	         "userName": "christophe"
	      }
	   },
	   {
	      "awsRegion": "sagov-westsouth-1r",
	      "eventCategory": "Management",
	      "eventID": "6e3e5c56-66d8-4e23-9a89-8498651357d5",
	      "eventName": "DescribeInstanceInformation",
	      "eventSource": "ssm.amazonaws.com",
	      "eventTime": "2024-08-02T09:08:20Z",
	      "eventType": "AwsApiCall",
	      "eventVersion": "1.08",
	      "managementEvent": true,
	      "readOnly": true,
	      "recipientAccountId": "056392974792",
	      "requestID": "8c004773-45de-49ee-aab8-44a83effbfd6",
	      "requestParameters": {
	         "filters": [
	            {
	               "key": "InstanceIds",
	               "values": [
	                  "i-00456A8D163f546Ff",
	                  "i-cfE23b1a7ceba6f86",
	                  "i-9D40CCFc0aE91CFa5"
	               ]
	            }
	         ]
	      },
	      "responseElements": null,
	      "sourceIPAddress": "253.252.51.07",
	      "tlsDetails": {
	         "cipherSuite": "ECDHE-RSA-AES128-GCM-SHA256",
	         "clientProvidedHostHeader": "ssm.sagov-westsouth-1r.amazonaws.com",
	         "tlsVersion": "TLSv1.2"
	      },
	      "userAgent": "stratus-red-team_ea782787-a65d-4fc4-9fca-1c97869a9a25",
	      "userIdentity": {
	         "accessKeyId": "AKIAW9X2Q2U25SK79UCX",
	         "accountId": "056392974792",
	         "arn": "arn:aws:iam::056392974792:user/christophe",
	         "principalId": "AIDA10CZIPPG73T21TDI",
	         "type": "IAMUser",
	         "userName": "christophe"
	      }
	   }
	]
    ```

[^1]: These logs have been gathered from a real detonation of this technique in a test environment using [Grimoire](https://github.com/DataDog/grimoire), and anonymized using [LogLicker](https://github.com/Permiso-io-tools/LogLicker).
