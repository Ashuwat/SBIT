import asyncio
import websockets

async def connect():
    uri = "ws://localhost:8080/ws"
    while True:
        try:
            async with websockets.connect(uri) as websocket:
                print("Connected to WebSocket server")
                async for message in websocket:
                    print(f"Received: {message}")
        except websockets.exceptions.ConnectionClosed:
            print("WebSocket connection closed, retrying...")
        except Exception as e:
            print(f"WebSocket error: {e}")
        
        await asyncio.sleep(1)  # Reconnect after 1 second

async def send_message():
    uri = "ws://localhost:8080/ws"
    async with websockets.connect(uri) as websocket:
        while True:
            message = input("Enter message: ")
            await websocket.send(message)

asyncio.run(connect())
