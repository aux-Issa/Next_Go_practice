import '../styles/globals.css'
import type { AppProps } from 'next/app'
import { QueryClient, QueryClientProvider } from 'react-query'
// import { QueryClient, QueryClientProvider } from '@tanstack/react-query'
import Task from '../components/task'

function MyApp({ Component, pageProps }: AppProps) {
  const queryClient = new QueryClient()
  return (
    <QueryClientProvider client={queryClient}>
      <Task/>
    </QueryClientProvider>
  )

}

export default MyApp
