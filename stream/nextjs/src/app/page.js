'use client'
// import Image from "next/image";
import { useEffect, useState } from "react";
import axios from "axios";

export default function Home() {
    // const [posts, setPosts] = useState(null);
    // const [loading, setLoading] = useState(true); 
    // const [error, setError] = useState(null);
    const [data, setData] = useState([]);

    // useEffect(() => {
    //     // const fetchPosts = async () => {
    //     //     try {
    //     //         // const response = await axios.get("/stream/stream-with-sse");
    //     //         const response = await axios.get("/stream/Z");
    //     //         console.log("response:", response);
    //     //         // setPosts(response.data);
    //     //     } catch (err) {
    //     //         // setError("Failed to fetch posts");
    //     //         console.log("err:", err);
    //     //     } finally {
    //     //         // setLoading(false);
    //     //         console.log("finnaly");
    //     //     }
    //     // };
    //     // fetchPosts()

    //     const fetchStream = async () => {
    //         try {
    //             const response = await fetch("/stream/stream-with-sse"); // Panggil endpoint backend
      
    //             if (!response.body) throw new Error("Response body is empty");
      
    //             const reader = response.body.getReader();
    //             const decoder = new TextDecoder();
    //             let receivedData = [];

    //             let i = 0
    //             while (true) {
    //                 const { value, done } = await reader.read();
    //                 // console.log(i++);
    //                 // console.log(value);
    //                 // console.log();
                    
                    
    //                 if (done) break;
      
    //                 const chunk = decoder.decode(value, { stream: true });
    //                 console.log(i++);
    //                 console.log("chunk:", chunk);
    //                 console.log();
    //                 receivedData.push(chunk);
    //                 setData([...receivedData]); // Perbarui state setiap kali ada data baru
    //             }

    //             // const reader = response.body.pipeThrough(new TextDecoderStream()).getReader();
    //             // let receivedData = [];

    //             // while (true) {
    //             //     const { value, done } = await reader.read();
    //             //     if (done) break; // Stop jika streaming selesai

    //             //     console.log("Chunk received:", value); // Log setiap chunk diterima

    //             //     receivedData.push(value);
    //             //     setData([...receivedData]); // Update UI langsung
    //             // }

    //             // while (true) {
    //             //     const reader = response.body.getReader();
    //             //     const decoder = new TextDecoder();
    //             //     let receivedData = [];
    //             //     const { value, done } = await reader.read();
    //             //     if (done) break;
    //             //     const chunk = decoder.decode(value, { stream: true });
    //             //     console.log("chunk:", chunk);
    //             //     receivedData.push(chunk);
    //             //     setData([...receivedData]);
    //             // }

    //             // while(true) {
    //             //     const reader = response.body.pipeThrough(new TextDecoderStream()).getReader();
    //             //     let receivedData = [];
    //             //     const { value, done } = await reader.read();
    //             //     if (done) break;
    //             //     console.log("Chunk received:", value);
    //             //     receivedData.push(value);
    //             //     setData([...receivedData]);
    //             // }
    //         } catch (error) {
    //             console.error("Streaming error:", error);
    //         }
    //     };
    //     fetchStream();
    // },[])

    //     useEffect(() => {
    //     const fetchStream = async () => {
    //       try {
    //         const response = await fetch("/stream/stream-with-sse"); // Panggil backend
    
    //         if (!response.body) throw new Error("Response body is empty");
    
    //         const reader = response.body.getReader();
    //         const decoder = new TextDecoder();
    //         let receivedData = [];
    
    //         const processStream = async () => {
    //           while (true) {
    //             const { value, done } = await reader.read();
    //             if (done) break; // Hentikan loop jika selesai
    
    //             const chunk = decoder.decode(value, { stream: true });
    
    //             // console.log("Chunk received:", chunk); // Log setiap chunk yang diterima
    //             console.log("mantap");
    //             console.log(chunk);
                
                
    
    //             receivedData.push(chunk);
    //             setData([...receivedData]); // Update UI setiap kali menerima data baru
    //           }
    //         };
    
    //         processStream();
    //       } catch (error) {
    //         console.error("Streaming error:", error);
    //       }
    //     };
    
    //     fetchStream();
    //   }, []);

    // useEffect(() => {

    // }, []);

    // useEffect(() => {
    //     const eventSource = new EventSource("/stream/stream-with-sse");
    //     // const eventSource = new EventSource("http://localhost:8080/stream/stream-with-sse");
    
    //     eventSource.onmessage = (event) => {
    //       console.log("SSE Data:", event.data);
    //       if (event.data.includes("error")) {
    //         console.error("Server Error:", event.data);
    //       }
    //       setData((prevData) => [...prevData, event.data]);
    //     };
    
    //     eventSource.onerror = (error) => {
    //       console.error("SSE Error:", error);
    //       eventSource.close();
    //       console.log("ReadyState:", eventSource.readyState);
    //     };
    
    //     return () => eventSource.close();
    // }, []);

    // useEffect(() => {
    //     const fetchPosts = async () => {
    //         try {
    //             // const response = await axios.get("/stream/stream-with-sse");
    //             const response = await axios.get("http://localhost:8080/stream/response");
    //             console.log("response:", response);
    //             // setPosts(response.data);
    //         } catch (err) {
    //             // setError("Failed to fetch posts");
    //             console.log("err:", err);
    //         } finally {
    //             // setLoading(false);
    //             console.log("finnaly");
    //         }
    //     };
    //     fetchPosts()
    // }, []);

    // useEffect(() => {
    //     const fetchStream = async () => {
    //       try {
    //         const response = await fetch("/stream/stream-without-channel"); // Panggil endpoint backend
    
    //         if (!response.body) throw new Error("Response body is empty");
    
    //         const reader = response.body.getReader();
    //         const decoder = new TextDecoder();
    //         let receivedData = [];
    
    //         while (true) {
    //           const { value, done } = await reader.read();
    //           if (done) break;
    
    //           const chunk = decoder.decode(value, { stream: true });
    //           console.log("mantap");
    //           console.log(chunk);
    //           console.log();
    //           receivedData.push(chunk);
    //           setData([...receivedData]); // Perbarui state setiap kali ada data baru
    //         }
    //       } catch (error) {
    //         console.error("Streaming error:", error);
    //       }
    //     };
    
    //     fetchStream();
    // }, []);

    // useEffect(() => {
    //     const fetchStream = async () => {
    //         try {
    //             const response = await fetch("/stream/stream-without-channel"); // Panggil endpoint backend
    
    //             if (!response.body) throw new Error("Response body is empty");

    //             while (true) {
    //                 const reader = response.body.getReader();
    //                 const decoder = new TextDecoder();
    //                 let receivedData = [];
    
    //                 // while (true) {
    //                 const { value, done } = await reader.read();
    //                 if (done) break;
    
    //                 const chunk = decoder.decode(value, { stream: true });
    //                 console.log("mantap");
    //                 console.log(chunk);
    //                 console.log();
    //                 receivedData.push(chunk);
    //                 setData([...receivedData]); // Perbarui state setiap kali ada data baru
    //             }
    //         } catch (error) {
    //             console.error("Streaming error:", error);
    //         }
    //     };
    
    //     fetchStream();
    // }, []);

    // useEffect(() => {
    //     const fetchStream = async () => {
    //         try {
    //             const response = await fetch("/stream/stream-without-channel"); // Panggil backend
            
    //             if (!response.body) throw new Error("Response body is empty");
    
    //             const reader = response.body.pipeThrough(new TextDecoderStream()).getReader();
    //             let receivedData = [];
    
    //             while (true) {
    //                 const { value, done } = await reader.read();
    //                 if (done) break; // Stop jika streaming selesai
    
    //                 console.log("Chunk received:", value); // Log setiap chunk diterima
    
    //                 receivedData.push(value);
    //                 setData([...receivedData]); // Update UI langsung
    //             }
    //         } catch (error) {
    //             console.error("Streaming error:", error);
    //         }
    //     };
    
    //     fetchStream();
    // }, []);

    // useEffect(() => {
    //     const fetchStream = async () => {
    //         try {
    //             const response = await fetch("/stream/stream-without-channel"); // Panggil backend
            
    //             if (!response.body) throw new Error("Response body is empty");

    //             while (true) {

    //                 const reader = response.body.pipeThrough(new TextDecoderStream()).getReader();
    //                 let receivedData = [];
    
    //             // while (true) {
    //                 const { value, done } = await reader.read();
    //                 if (done) break; // Stop jika streaming selesai

    //                 console.log("mantap");
    //                 console.log("Chunk received:", value); // Log setiap chunk diterima
    //                 console.log();
    
    //                 receivedData.push(value);
    //                 setData([...receivedData]); // Update UI langsung
    //             }
    //         } catch (error) {
    //             console.error("Streaming error:", error);
    //         }
    //     };
    
    //     fetchStream();
    // }, []);

    // useEffect(() => {
    //     const fetchStream = async () => {
    //         try {
    //             const response = await fetch("http://localhost:8080/stream/stream-without-channel");
    //             const reader = response.body.getReader();
    //             const decoder = new TextDecoder();

    //             while (true) {
    //                 const { done, value } = await reader.read();
    //                 if (done) break;
    //                 // const chunk = decoder.decode(value, { stream: true });
    //                 console.log(decoder.decode(value));
    //                 // output.innerHTML += `<p>${decoder.decode(value)}</p>`;
    //             }
    //         } catch (error) {
    //             // output.innerHTML += `<p style="color:red;">Error: ${error.message}</p>`;
    //             console.log("error:", error);
    //         }
    //     }

    //     fetchStream()
    // }, [])

    useEffect(() => {
        const fetchStream = async () => {
            try {
                // i dont know why have to put http://localhost:8080 to make stream works, or create middleware
                // const response = await fetch("http://localhost:8080/stream/stream-without-channel"); // Panggil endpoint backend
                const response = await fetch("/stream/stream-without-channel");
                if (!response.body) throw new Error("Response body is empty");
                const reader = response.body.getReader();
                const decoder = new TextDecoder();
                // let receivedData = [];
    
                while (true) {
                    const { value, done } = await reader.read();
                    if (done) break;
    
                    const chunk = decoder.decode(value, { stream: true });
                    console.log("chunk:", chunk);
                    
                    // receivedData.push(chunk);
                    // setData([...receivedData]); // Perbarui state setiap kali ada data baru
                }
            } catch (error) {
                console.error("Streaming error:", error);
            }
        };
    
        fetchStream();
    }, []);

    return (
        <>
            {/* <div>
                <h1>Streaming Data</h1>
                <ul>
                    {data.map((item, index) => (
                        <li key={index}>{item}</li>
                    ))}
                </ul>
            </div> */}
            {/* <div>
                <h1>Server-Sent Events (SSE) Streaming</h1>
                <ul>
                    {data.map((item, index) => (
                    <li key={index}>{item}</li>
                    ))}
                </ul>
            </div> */}
        </>
    );
}
