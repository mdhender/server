import {configureStore} from '@reduxjs/toolkit';
import siteReducer from '../features/site/siteSlice';

export default configureStore({
	reducer: {
		site: siteReducer,
	},
});
