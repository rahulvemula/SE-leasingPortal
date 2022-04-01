import { createSlice } from '@reduxjs/toolkit'

const authSlice = createSlice({
  name: 'isAuthenticated',
  initialState: { value: false, usertoken: '' },
  reducers: {
    loginSuccessful: (state, action) => {
      state.usertoken = action.payload.token;
      state.value = true;
    },
    logoutSuccessful: state => {
      state.value = false;
      state.usertoken = '';
    }
  }
})

export const { loginSuccessful, logoutSuccessful } = authSlice.actions;
export default authSlice.reducer;