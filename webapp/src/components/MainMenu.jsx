import React from 'react';
import {Link} from "@reach/router";
import {useSelector} from "react-redux";

function MainMenu() {
	const items = useSelector(state => state.site.menus.nav);

	return (
			<div id="mainmenu">
				{items.map((item, key) => {
					return (
							<dl key={key} className={item.class}>
								<dt><Link to={item.link}>{item.label}</Link></dt>
								{item.children.map((child, key) => {
									if (child.link) {
										return <dd key={key}><Link to={child.link}>{child.label}</Link></dd>;
									}
									return <dd key={key}>{child.label}</dd>;
								})}
							</dl>);
				})}
				<hr className="clear" />
			</div>);
}

export default MainMenu;
