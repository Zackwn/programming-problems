const addTwoNumbers = (l1, l2) => {
  let head = null
  let tail = null

  let prevTenth = 0
  while (l1 || l2) {
    let sum = 0

    sum += l1 ? l1.val : 0
    sum += l2 ? l2.val : 0

    l1 = l1 ? l1.next : l1
    l2 = l2 ? l2.next : l2

    const currentValue = (prevTenth + sum) % 10 // take last digit
    prevTenth = Math.floor((prevTenth + sum) / 10) // take the tenth onwards

    if (!head) {
      head = new ListNode(currentValue)
      tail = head
    } else {
      tail.next = new ListNode(currentValue)
      tail = tail.next
    }
  }

  // if there is prevTenth after the iteration we need to add to our list
  if (prevTenth > 0) {
    tail.next = new ListNode(prevTenth)
  }

  return head
}