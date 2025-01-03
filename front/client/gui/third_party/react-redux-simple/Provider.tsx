import React from 'react'
import {ReactReduxContext} from './Context'

export default function Provider(providerProps) {
  const { children, store } = providerProps

  const contextValue = React.useMemo(() => ({store}), [store])

  return (
    <ReactReduxContext.Provider value={contextValue}>
      {children}
    </ReactReduxContext.Provider>
  )
}