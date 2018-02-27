// Get all the  buttons from the page
let a = document.querySelectorAll('.ruResponseButtons');

// Convert the node list into an array
a = [...a]

a.forEach(buttons => {
	buttons.childNodes[1].click()
})
