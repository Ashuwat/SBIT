async function fetchData() {
    try {
        const response = await fetch("http://localhost:3030/data");
        if (!response.ok) throw new Error("Network error");

        const data = await response.json(); 
        console.log("Received data:", data);

        const outputDiv = document.getElementById("output");
        outputDiv.innerHTML = "<h3>Simulation Data:</h3>";

        data.forEach((row, index) => {
            outputDiv.innerHTML += `<p>Node ${index + 1}: [${row[0]}, ${row[1]}, ${row[2]}]</p>`;
        }); 

    } catch (error) {
        console.error("Error fetching data:", error);
        document.getElementById("output").innerHTML = "<p style='color: red;'>Error fetching data.</p>";
    }
}

document.getElementById("fetchDataBtn").addEventListener("click", fetchData);
