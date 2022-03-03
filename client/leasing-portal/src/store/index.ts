import { configureStore } from '@reduxjs/toolkit'
import authReducer from './auth'
import userDataReducer from './userData'

export const store = configureStore({
  reducer: {
    isAuthenticated: authReducer,
    userData: userDataReducer
  },
})