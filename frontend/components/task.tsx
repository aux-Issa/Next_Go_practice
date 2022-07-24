import {useQuery} from 'react-query';

export default function Task() {
  const {isLoading, isError, data} = useQuery('repodata', 
    async () => {
      const res = await fetch('http://localhost:1323/tasks')
      const data = await res.json()
    }
  )
  if (isLoading) {
    return <p>Loading...</p>
  }
  if (isError) {
    return <p>Error!</p>
  }
  return (
    <div className='flex min-h-screen flex-col items-center justify-center font-mono text-gray-8080'>
      <h1>TODOリスト(一覧取得)</h1>
      <ul>
        {data.map((task:any) => (
          <li key={task.ID}>
            {task.Title}: {task.Content}
          </li>
        ))}
      </ul>
    </div>
  )
}