import { createSlice } from '@reduxjs/toolkit'

const userSlice = createSlice({
  name: 'userData',
  initialState: { username: '', email:'', name:'' },
  reducers: {
    updateUserData: (state, action) => {
      let data = action.payload;
      state.username = data.username;
      state.email = data.email;
      state.name = data.name;
    },
    resetUserData: state => {
      state.username ='';
      state.email = '';
      state.name = '';
    }
  }
})

export const { updateUserData, resetUserData } = userSlice.actions;
export default userSlice.reducer;