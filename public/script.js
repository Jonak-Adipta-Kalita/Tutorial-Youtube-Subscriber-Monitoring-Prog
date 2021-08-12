let subscribers = {};
const websocket = new WebSocket("ws://localhost:8000/stats");

websocket.onopen = (event) => console.log("Successfully Connected to Websocket!!");
websocket.onerror = (err) => alert(err);
websocket.onmessage = (event) => {
	subscribers = JSON.parse(event.data);
	console.log(subscribers);
	document.getElementById("subs").innerText = subscribers.statistics.subscriberCount;
};
