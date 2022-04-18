import { configureStore } from '@reduxjs/toolkit'
import authReducer from './auth'
import userDataReducer from './userData'
import leaseDataReducer from './leaseData';

export const store = configureStore({
  reducer: {
    isAuthenticated: authReducer,
    userData: userDataReducer,
    leaseData: leaseDataReducer,
  },
})