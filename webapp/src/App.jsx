import React, {useEffect} from 'react';
import {useDispatch, useSelector} from 'react-redux';
import {Router} from "@reach/router";

import Banner from "./components/Banner";
import Footer from "./components/Footer";
import Home from "./components/Home";
import MainMenu from "./components/MainMenu";
import Sidebar from "./components/Sidebar";

function App() {
	console.log('app');

	const dispatch = useDispatch();
	const site = useSelector(state => state.site);

	return (<>
		<Banner/>
		<div id="wrap">
			<MainMenu/>
			<div id="content">
				<Router>
					<Home path="/"/>
				</Router>
			</div>
			<Sidebar/>
			<hr className={"clear"}/>
		</div>
		<Footer/>
	</>);
}

export default App;
