import type { NextPage } from 'next'
import Alert from '../../../Common/Alert'
import { AddEdit } from 'Common/Products'
const Add: NextPage = (params) => {
  return (
     <div>
      <Alert /> 
      <AddEdit params={params} />
     </div> 
  )
}

export default Add 