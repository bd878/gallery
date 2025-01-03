import React from 'react'
import hoistStatics from './hoistStatics';
import bindActionCreators from './bindActionCreators';
import {ReactReduxContext} from './Context'

function is(x: unknown, y: unknown) {
  if (x === y) {
    return x !== 0 || y !== 0 || 1 / x === 1 / y
  } else {
    return x !== x && y !== y
  }
}

function strictEqual(objA: any, objB: any) {
  if (is(objA, objB)) return true

  if (
    typeof objA !== 'object' ||
    objA === null ||
    typeof objB !== 'object' ||
    objB === null
  ) {
    return false
  }

  const keysA = Object.keys(objA)
  const keysB = Object.keys(objB)

  if (keysA.length !== keysB.length) return false

  for (let i = 0; i < keysA.length; i++) {
    if (
      !Object.prototype.hasOwnProperty.call(objB, keysA[i]) ||
      !is(objA[keysA[i]], objB[keysA[i]])
    ) {
      return false
    }
  }

  return true
}


function shallowEqual(a, b) {
  return a === b
}

function mergeProps(stateProps, dispatchProps, ownProps) {
  return {...stateProps, ...dispatchProps, ...ownProps}
}

export function connect(mapStateToProps, mapDispatchToProps) {
  function wrapWithConnect(WrappedComponent) {
    const wrappedComponentName =
      WrappedComponent.displayName || WrappedComponent.name || 'Component';

    const displayName = `Connect(${wrappedComponentName})`

    function ConnectFunction(props) {
      const contextValue = React.useContext(ReactReduxContext)
      const store = contextValue.store

      const makePropsSnapshot = React.useMemo(() => {
        const selector = () => {
          const stateProps = mapStateToProps(store.getState())
          const dispatchProps = bindActionCreators(mapDispatchToProps, store.dispatch)
          // TODO: cache last prop values; otherwise useSyncExternalStore falls to infinite loop
          return mergeProps(stateProps, dispatchProps, props)
        }

        return selector
      }, [
        store,
        mapStateToProps,
        mapDispatchToProps,
        props,
      ])

      let childProps = {}
      try {
        childProps = React.useSyncExternalStore(store.subscribe, makePropsSnapshot)
      } catch (e) {
        console.error(e)
        throw e
      }

      const renderedWrappedComponent = React.useMemo(() => {
        return (
          <WrappedComponent {...childProps} />
        )
      }, [childProps])

      return renderedWrappedComponent
    }

    const Connect = React.memo(ConnectFunction)
    Connect.WrappedComponent = WrappedComponent
    Connect.displayName = ConnectFunction.displayName = displayName

    return hoistStatics(Connect, WrappedComponent)
  }

  return wrapWithConnect
}