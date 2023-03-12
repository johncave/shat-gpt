import './Log.css'
import useWebSocket from 'react-use-websocket';
import { useEffect, useState } from 'react';

const getWsURL = () => {
  console.log('getting url in dev')
  if (import.meta.env.MODE === 'development') {
    return 'ws://localhost:8080/ws/shatgpt'
  }
  console.log('getting url in prod')
  return `ws://${document.location.host}/ws/shatgpt`
}

function Log() {

  const [socketUrl, setSocketUrl] = useState('wss://echo.websocket.org');
  const [messageHistory, setMessageHistory] = useState([]);

  const { sendMessage, sendJsonMessage, lastMessage, readyState } = useWebSocket(getWsURL);

  useEffect(() => {
    if (lastMessage !== null) {
      setMessageHistory((prev) => prev.concat(lastMessage));
    }
  }, [lastMessage, setMessageHistory]);

console.log('message History', messageHistory)

  // const {
  //   sendMessage,
  //   sendJsonMessage,
  //   lastMessage,
  //   lastJsonMessage,
  //   readyState,
  //   getWebSocket,
  // } = useWebSocket(getWsURL, {
  //   onOpen: () => {
  //     console.log("Opened")
  //   },
  //   onMessage: (m) => {
  //     console.log({m, messageReceived: "dsfsdsf"})
  //   },
  //   shouldReconnect: (closeEvent) => true,
  // });

  const handleClick = () => {
    console.log('💩')
    sendMessage('vanilla message')
    sendJsonMessage({
      eventName: 'yolo',
      message: '💩'
    })
  }

  return (
    <div className="board">
      <div id="log" className='log'></div>
      <div className="poop-button__container">
        <button onClick={ handleClick } className='poop-button'>💩</button>
      </div>
    </div>
  )
}

export default Log
