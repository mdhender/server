import React from 'react';
import {Link} from "@reach/router";

const SideMenu = ({side, title, items, children}) => {
	return (
			<div className={side}>
				<h2>{title}</h2>
				<ul className="sidemenu">
					{items.map((item, key) => {
						return (
								<li key={key}>
									<Link to={item.link}>{item.label}</Link>
									{!item.children ? <></> :
											<ul>
												{item.children.map((child, key) => {
													return <li key={key}><Link to={child.link}>{child.label}</Link></li>;
												})}
											</ul>
									}
								</li>);
					})}
				</ul>
				{children}
			</div>
	);
};

export default SideMenu;