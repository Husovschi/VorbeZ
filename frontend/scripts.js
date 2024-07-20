document.addEventListener('DOMContentLoaded', () => {
     const ws = new WebSocket('wss://${window.location.host}/ws');
    const chatWindow = document.getElementById('chat-window');
    const messageInput = document.getElementById('message-input');
    const sendButton = document.getElementById('send-button');
  
    ws.onmessage = (event) => {
        const message = JSON.parse(event.data);
        displayMessage(message.content, 'received');
    };
  
    sendButton.addEventListener('click', () => {
        const message = messageInput.value;
        if (message.trim() !== '') {
            ws.send(JSON.stringify({ content: message }));
            displayMessage(message, 'sent');  // Display sent message
            messageInput.value = '';
        }
    });
  
    messageInput.addEventListener('keyup', (event) => {
        if (event.key === 'Enter') {
            sendButton.click();
        }
    });
  
    function displayMessage(message, type) {
        const messageElement = document.createElement('div');
        messageElement.textContent = message;
        messageElement.classList.add('message', type);
        chatWindow.appendChild(messageElement);
        chatWindow.scrollTop = chatWindow.scrollHeight;
    }
  });
  