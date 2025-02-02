<script setup>
    import { ref, onMounted } from 'vue'

    onMounted(async () => {
        // const eventSource = new EventSource("http://localhost:8080/sse/handle-sse-without-channel");
        const eventSource = new EventSource("/sse/handle-sse-without-channel");
        eventSource.onmessage = (event) => {
            console.log("SSE Data:", event.data);
            console.log("event:", event);
        }

        eventSource.onerror = (error) => {
            console.log("error:", error);
            eventSource.close();
        }
    })
</script>

<template>
  <div>
    <NuxtRouteAnnouncer />
    <NuxtWelcome />
  </div>
</template>
