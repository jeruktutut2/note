'use client'

import Image from "next/image";
import { useEffect, useRef, useState } from "react";

export default function Home() {
  const [clientId, setClientId] = useState("");
  const [clientIdSendTo, setClientIdSendTo] = useState("");
  const [message, setMessage] = useState("");
  const [messages, setMessages] = useState([]);
  const ws = useRef(null);

  function connectWebSocket() {
    if (!clientId) {
      alert("Masukkan ID dulu!");
      return;
    }

    if (ws.current) {
      ws.current.close();
    }

    ws.current = new WebSocket(`ws://localhost:8080/ws?id=${clientId}`);

    ws.current.onopen = () => {
      console.log(`Terhubung sebagai ${clientId}`);
    };

    ws.current.onmessage = (event) => {
      setMessages((prevMessages) => [...prevMessages, event.data]);
    };

    ws.current.onclose = () => {
      console.log("WebSocket terputus");
    };
  };

  async function sendMessage() {
    if (!clientId) {
      alert("Hubungkan WebSocket dulu!");
      return;
    }

    if (!message) {
      alert("Masukkan pesan!");
      return;
    }

    await fetch(`http://localhost:8080/send-message?clientIdSendTo=${clientIdSendTo}&msg=${message}`);
    setMessages((prevMessages) => [...prevMessages, `Anda: ${message}`]);
    setMessage("");
  };

  useEffect(() => {
    return () => {
      if (ws.current) {
        ws.current.close();
      }
    };
  }, []);
  return (
    // <div className="grid grid-rows-[20px_1fr_20px] items-center justify-items-center min-h-screen p-8 pb-20 gap-16 sm:p-20 font-[family-name:var(--font-geist-sans)]">
    //   <main className="flex flex-col gap-8 row-start-2 items-center sm:items-start">
    //     <Image
    //       className="dark:invert"
    //       src="/next.svg"
    //       alt="Next.js logo"
    //       width={180}
    //       height={38}
    //       priority
    //     />
    //     <ol className="list-inside list-decimal text-sm text-center sm:text-left font-[family-name:var(--font-geist-mono)]">
    //       <li className="mb-2">
    //         Get started by editing{" "}
    //         <code className="bg-black/[.05] dark:bg-white/[.06] px-1 py-0.5 rounded font-semibold">
    //           src/app/page.js
    //         </code>
    //         .
    //       </li>
    //       <li>Save and see your changes instantly.</li>
    //     </ol>

    //     <div className="flex gap-4 items-center flex-col sm:flex-row">
    //       <a
    //         className="rounded-full border border-solid border-transparent transition-colors flex items-center justify-center bg-foreground text-background gap-2 hover:bg-[#383838] dark:hover:bg-[#ccc] text-sm sm:text-base h-10 sm:h-12 px-4 sm:px-5"
    //         href="https://vercel.com/new?utm_source=create-next-app&utm_medium=appdir-template-tw&utm_campaign=create-next-app"
    //         target="_blank"
    //         rel="noopener noreferrer"
    //       >
    //         <Image
    //           className="dark:invert"
    //           src="/vercel.svg"
    //           alt="Vercel logomark"
    //           width={20}
    //           height={20}
    //         />
    //         Deploy now
    //       </a>
    //       <a
    //         className="rounded-full border border-solid border-black/[.08] dark:border-white/[.145] transition-colors flex items-center justify-center hover:bg-[#f2f2f2] dark:hover:bg-[#1a1a1a] hover:border-transparent text-sm sm:text-base h-10 sm:h-12 px-4 sm:px-5 sm:min-w-44"
    //         href="https://nextjs.org/docs?utm_source=create-next-app&utm_medium=appdir-template-tw&utm_campaign=create-next-app"
    //         target="_blank"
    //         rel="noopener noreferrer"
    //       >
    //         Read our docs
    //       </a>
    //     </div>
    //   </main>
    //   <footer className="row-start-3 flex gap-6 flex-wrap items-center justify-center">
    //     <a
    //       className="flex items-center gap-2 hover:underline hover:underline-offset-4"
    //       href="https://nextjs.org/learn?utm_source=create-next-app&utm_medium=appdir-template-tw&utm_campaign=create-next-app"
    //       target="_blank"
    //       rel="noopener noreferrer"
    //     >
    //       <Image
    //         aria-hidden
    //         src="/file.svg"
    //         alt="File icon"
    //         width={16}
    //         height={16}
    //       />
    //       Learn
    //     </a>
    //     <a
    //       className="flex items-center gap-2 hover:underline hover:underline-offset-4"
    //       href="https://vercel.com/templates?framework=next.js&utm_source=create-next-app&utm_medium=appdir-template-tw&utm_campaign=create-next-app"
    //       target="_blank"
    //       rel="noopener noreferrer"
    //     >
    //       <Image
    //         aria-hidden
    //         src="/window.svg"
    //         alt="Window icon"
    //         width={16}
    //         height={16}
    //       />
    //       Examples
    //     </a>
    //     <a
    //       className="flex items-center gap-2 hover:underline hover:underline-offset-4"
    //       href="https://nextjs.org?utm_source=create-next-app&utm_medium=appdir-template-tw&utm_campaign=create-next-app"
    //       target="_blank"
    //       rel="noopener noreferrer"
    //     >
    //       <Image
    //         aria-hidden
    //         src="/globe.svg"
    //         alt="Globe icon"
    //         width={16}
    //         height={16}
    //       />
    //       Go to nextjs.org →
    //     </a>
    //   </footer>
    // </div>

    <>
      <h1>WebSocket Client (Next.js)</h1>

      <input type="text" placeholder="Masukkan ID Anda" value={clientId} onChange={(e) => setClientId(e.target.value)} style={{ padding: "10px", margin: "5px", width: "80%" }}/>
      <button onClick={connectWebSocket} style={{ padding: "10px", cursor: "pointer" }}>Hubungkan</button>

      <br />
      <input type="text" placeholder="Masukkan ID Send To Anda" value={clientIdSendTo} onChange={(e) => setClientIdSendTo(e.target.value)} style={{ padding: "10px", margin: "5px", width: "80%" }}/>

      <br />

      <input type="text" placeholder="Ketik pesan..." value={message} onChange={(e) => setMessage(e.target.value)} style={{ padding: "10px", margin: "5px", width: "80%" }}/>
      <button onClick={sendMessage} style={{ padding: "10px", cursor: "pointer" }}>Kirim</button>

      <div style={{ marginTop: "20px", textAlign: "left" }}>
        <h3>Pesan:</h3>
        <ul>
          {messages.map((msg, index) => (
            <li key={index}>{msg}</li>
          ))}
        </ul>
      </div>

    </>
  );
}
