function cakes(recipe, available) {
  let cakeCount = 0 
  while (true) {
    let ok = true
    for (let key in recipe) {
      if (available[key] - recipe[key] >= 0) {
        available[key] = available[key] - recipe[key]
      } else {
        ok = false
        break
      }
    }
    if (!ok) {
      break
    }
    cakeCount++
  }
  return cakeCount
}