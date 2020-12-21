import React from 'react';
import {useSelector} from 'react-redux';
import {selectSlug, selectTitle,} from '../features/site/siteSlice';
import {Link} from "@reach/router";

function Banner() {
	const title = useSelector(selectTitle);
	const slug = useSelector(selectSlug);
	return (
			<div id="top">
				<p id="skiplinks">Skip to: <a href="#content">content</a> | <a href="#sidebar">sidebar</a></p>
				<div id="sitetitle">
					<h1><Link to="/">{title}</Link></h1>
					<p>{slug}</p>
				</div>
				<hr className="clear"/>
			</div>);
}

export default Banner;