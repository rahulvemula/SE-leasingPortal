import { createSlice } from '@reduxjs/toolkit'

const authSlice = createSlice({
  name: 'isAuthenticated',
  initialState: { value: false },
  reducers: {
    loginSuccessful: state => {
      state.value = true
    },
    logoutSuccessful: state => {
      state.value = false
    }
  }
})

export const { loginSuccessful, logoutSuccessful } = authSlice.actions;
export default authSlice.reducer;