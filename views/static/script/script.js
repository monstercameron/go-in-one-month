console.log("script.js loaded");

// wait for the DOM to be loaded
document.addEventListener("DOMContentLoaded", function (event) {
  // Your code here
  console.log("DOM fully loaded and parsed");
});


document.addEventListener('htmx:afterSwap', function(event) {
    // Check if the target is the #todos container
    if (event.detail.target.id === 'todos') {
        // Clear the input field
        document.getElementById('add-todo').value = '';
    }
});