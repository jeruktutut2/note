export default defineEventHandler(async (event) => {
    const config = useRuntimeConfig()
    // const userId = getRouterParam(event, "id")
  
    // Redirect ke stream yang baru
    const targetUrl = `${config.public.streamBase}/stream/stream-without-channel`
    return await proxyRequest(event, targetUrl)
})