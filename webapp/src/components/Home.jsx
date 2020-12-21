import React from 'react';
import {Link} from "@reach/router";

import Trio from "./Trio";

function Home() {
	console.log('home');
	return (<>
		<h2><Link to="/">Introducing: Gas Giant Battles</Link></h2>
		<p className="introtext">
			Gas Giant Battles is a strategic, turn-based game of conquest and accounting.
		</p>
		<p>
			Players begin the game in control of one star-system.
			Lectus nulla at volutpat diam ut venenatis tellus in.
			Accumsan tortor posuere ac ut.
			Accumsan tortor posuere ac ut consequat semper viverra nam.
			Feugiat nisl pretium fusce id velit ut tortor pretium viverra.
		</p>
		<p>
			Lorem donec massa sapien faucibus et molestie ac feugiat sed.
			Duis convallis convallis tellus id.
			Viverra nibh cras pulvinar mattis nunc sed blandit.
		</p>
		<p>
			Eu non diam phasellus vestibulum lorem sed risus ultricies tristique.
			Et tortor consequat id porta nibh.
			Gravida quis blandit turpis cursus in hac habitasse platea.
			Sed ullamcorper morbi tincidunt ornare massa eget.
		</p>
		<p>
			Consectetur adipiscing elit duis tristique sollicitudin nibh sit amet.
			Fusce id velit ut tortor.
		</p>
		<p>
			Read more...
		</p>

		<Trio column={"1"}>
			<h3><Link to="/support-the-development">Support the development</Link></h3>
			<p>
				If you want to support the development of this game, there are plenty of opportunities.
			</p>
			<ul>
				<li>Coding</li>
				<li>Testing</li>
				<li>Documenting</li>
				<li>Spreading the word</li>
			</ul>
			<p><Link to="/support-the-development">Read more...</Link></p>
		</Trio>

		<Trio column={"2"}>
			<h3><Link to="/how-to-join">How to join</Link></h3>
			<p>
				Libero volutpat sed cras ornare arcu dui.
				Accumsan lacus vel facilisis volutpat.
			</p>
			<p>
				Sit amet risus nullam eget felis eget nunc.
				Scelerisque eleifend donec pretium vulputate sapien nec.
				Placerat duis ultricies lacus sed turpis tincidunt id aliquet.
			</p>
			<p>
				In ornare quam viverra orci sagittis eu volutpat.
				Pellentesque habitant morbi tristique senectus et netus et.
			</p>
			<p><Link to="/how-to-join">Read more...</Link></p>
		</Trio>

		<Trio column={"3"}>
			<h3><Link to="/customization">Customization</Link></h3>
			<p>
				Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore
				magna aliqua.
				Egestas sed sed risus pretium quam. Ut aliquam purus sit amet luctus.
			</p>
			<p>
				Pellentesque habitant morbi tristique senectus et netus et malesuada.
				Integer quis auctor elit sed vulputate mi sit amet.
				Duis ultricies lacus sed turpis tincidunt id aliquet risus.
				Ut morbi tincidunt augue interdum velit euismod in pellentesque massa.
				A cras semper auctor neque vitae tempus quam.
			</p>
			<p><Link to="/customization">Read more...</Link></p>
		</Trio>
	</>);
}

export default Home;
