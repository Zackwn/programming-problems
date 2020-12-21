def passwordCracker(passwords, loginAttempt):
    index = 0
    while True:
        if index == len(passwords): return "WRONG PASSWORD"
        p = passwords[index]
        i = loginAttempt.find(p)
        if (i != -1):
            if i == 0:
                la = loginAttempt[len(p):]
                if len(la) == 0:
                    return p
                return p + " " + passwordCracker(passwords, la)
        index += 1


passwords1 = ["zpkovasi", "qzobpkmdlu", "ghevtwlypw", "lhaqspnhja", "vdgbg", "wdjrvql", "azratm", "vcaqnfvkbr", "ixochdkzhr", "ztnybiabk"]
loginAttempt1 = "ztnybiabkqzobpkmdluwdjrvqlazratmghevtwlypwazratmixochdkzhrwdjrvqlvdgbgztnybiabkghevtwlypwlhaqspnhjazpkovasiwdjrvqlixochdkzhrlhaqspnhjazpkovasivdgbglhaqspnhjaixochdkzhrqzobpkmdluazratmqzobpkmdluzpkovasiixochdkzhrvcaqnfvkbrvcaqnfvkbrqzobpkmdluqzobpkmdluazratmghevtwlypwvdgbgvcaqnfvkbrvdgbgghevtwlypwghevtwlypwvdgbgqzobpkmdluwdjrvqlazratmvcaqnfvkbrqzobpkmdluvdgbgghevtwlypwzpkovasilhaqspnhjaixochdkzhrghevtwlypwqzobpkmdluvdgbgqzobpkmdlulhaqspnhjalhaqspnhjaixochdkzhrwdjrvqlazratmlhaqspnhjazpkovasighevtwlypwghevtwlypwwdjrvqlzpkovasivdgbgqzobpkmdluazratmlhaqspnhjawdjrvqlvcaqnfvkbrqzobpkmdluqzobpkmdlughevtwlypwixochdkzhrixochdkzhrwdjrvqlwdjrvqlztnybiabkazratmlhaqspnhjaazratmixochdkzhrzpkovasivcaqnfvkbrazratmixochdkzhrghevtwlypwvcaqnfvkbrazratmvdgbgwdjrvqlazratmwdjrvqlghevtwlypwqzobpkmdluvcaqnfvkbrzpkovasilhaqspnhjaqzobpkmdlughevtwlypwwdjrvqlvcaqnfvkbrixochdkzhrghevtwlypwghevtwlypwixochdkzhrazratmghevtwlypwvdgbgzpkovasiazratmztnybiabkvcaqnfvkbrztnybiabkvcaqnfvkbrvdgbgqzobpkmdluvdgbgwdjrvqlvdgbg"

passwords2 = ["a", "aa", "aaa", "aaaa", "aaaaa", "aaaaaa", "aaaaaaa", "aaaaaaaa", "aaaaaaaaa", "aaaaaaaaaa"]
loginAttempt2 = "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaab"

result1 = passwordCracker(passwords1, loginAttempt1)
result2 = passwordCracker(passwords2, loginAttempt2)

for result in range([result1, result2]):
  if "WRONG PASSWORD" in result:
    print("WRONG PASSWORD")
  else:
    print(result)