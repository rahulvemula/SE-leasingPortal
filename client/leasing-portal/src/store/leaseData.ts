import { createSlice } from '@reduxjs/toolkit'

const leaseSlice = createSlice({
  name: 'leaseData',
  initialState: { name: '', email: '', startDate: '' , endDate: ''},
  reducers: {
    updateLeaseData: (state, action) => {
      let data = action.payload;
      state.name = data.name;
      state.email = data.email;
      state.startDate = data.startDate;
      state.endDate = data.endDate;
    },
  }
})

export const { updateLeaseData } = leaseSlice.actions;
export default leaseSlice.reducer;