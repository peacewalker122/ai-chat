document.addEventListener("DOMContentLoaded", function () {
	document.getElementById("send").addEventListener("click", function () {
		const message = document.getElementById("message").value;
		sendMessage(message);
		document.getElementById("message").value = ""; // clear the input after send
	});

	const chatWindow = document.getElementById("chat");

	const ws = new WebSocket("ws://localhost:8080/v1/chat");

	ws.onopen = function () {
		console.log("Connected to the server");
	};

	ws.onmessage = function (event) {
		const data = JSON.parse(event.data); // Parse the JSON data
		if (data.Parts && Array.isArray(data.Parts)) {
			data.Parts.forEach((part) => displayMessage(part)); // Display each part separately
		}
	};

	ws.onclose = function () {
		console.log("Disconnected from the server");
	};

	function sendMessage(message) {
		if (!message) return;
		ws.send(message);
	}

	function displayMessage(message) {
		const messageElement = document.createElement("div");
		messageElement.innerHTML = "<md-block> " + message + " </md-block>";
		messageElement.className = "bg-white p-2 rounded shadow-sm";
		chatWindow.appendChild(messageElement);
		chatWindow.scrollTop = chatWindow.scrollHeight; // auto-scroll to latest message
	}
});
