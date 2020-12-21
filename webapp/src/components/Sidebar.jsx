import React from 'react';
import {Link} from "@reach/router";
import {useSelector} from "react-redux";

import SideMenu from "./SideMenu";

function Sidebar() {
	const items = useSelector(state => state.site.menus.side);
	const notice = useSelector(state => state.site.notice);

	return (
			<div id="sidebar">
				<h2>Eleifend pretium</h2>
				<p>Tellus id interdum velit laoreet id. Commodo quis tincidunt nunc pulvinar sapien et ligula.</p>
				<SideMenu side={"left"} title={"Games"} items={items.left}>
					<h2>Documentation</h2>
					<ul>
						<li><Link to="/rules">Rules</Link></li>
						<li><Link to="/submitting-orders">Submitting Orders</Link></li>
						<li><Link to="/tutorials">Tutorials</Link></li>
						<li><Link to="/known-issues">Known Issues</Link></li>
					</ul>
					<h2>Imperdiet Massa</h2>
					<p>
						Dolor sed viverra ipsum nunc aliquet bibendum enim facilisis gravida.
						Facilisis mauris sit amet massa vitae tortor condimentum lacinia.
					</p>
				</SideMenu>
				<SideMenu side={"right"} title={"Game 100, Turn 13"} items={items.right}>
					<h2>Sample Links</h2>
					<ul>
						{items.samples.map((item, key) => {
							return <li key={key}><a href={item.link}>{item.label}</a></li>;
						})}
					</ul>
				</SideMenu>
				<hr className={"clear"}/>
				<h2>{notice.title}</h2>
				<p>{notice.text}</p>
				<hr className={"clear"}/>
			</div>
	);
}

export default Sidebar;
