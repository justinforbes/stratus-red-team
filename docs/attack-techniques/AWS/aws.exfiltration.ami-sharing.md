---
title: Exfiltrate an AMI by Sharing It
---

# Exfiltrate an AMI by Sharing It 

Platform: AWS

## MITRE ATT&CK Tactics


- Exfiltration

## Description


Exfiltrates an AMI by sharing it with an external AWS account.

Warm-up: Create an AMI.

Detonation: Share the AMI.


## Instructions

```bash title="Detonate with Stratus Red Team"
stratus detonate aws.exfiltration.ami-sharing
```