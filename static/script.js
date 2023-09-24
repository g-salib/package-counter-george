document.addEventListener("DOMContentLoaded", function() {
    const calculateButton = document.getElementById("calculateButton");
    const orderQuantityInput = document.getElementById("orderQuantity");
    const output = document.getElementById("output");

    calculateButton.addEventListener("click", function() {
        const orderQuantity = parseFloat(orderQuantityInput.value);

        // Make an API request using Fetch with the orderQuantity
        fetch(`/api/v1/packs/calculate?quantity=${orderQuantity}`)
            .then(response => response.json())
            .then(data => {
                // Access the "result" array directly from the API response
                const results = data.result;

                // Display the results in the output area
                const resultText = results.join(', '); // Join results with a comma and space
                output.textContent = `Results: ${resultText}`;
            })
            .catch(error => {
                console.error("Error:", error);
            });
    });
});
