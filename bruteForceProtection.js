const IPs = new Map()

function bruteForceDetected(loginRequest)
{
  if (IPs.has(loginRequest.sourceIP) === false) {
    IPs.set(loginRequest.sourceIP, 0)
  }
  
  if (loginRequest.successful) {
    IPs.set(loginRequest.sourceIP, 0)
    return false
  } else {
    const timesFailed = IPs.get(loginRequest.sourceIP)
    if (timesFailed === 19) {
      return true
    }
    IPs.set(loginRequest.sourceIP, timesFailed+1)
  }
  return false
}