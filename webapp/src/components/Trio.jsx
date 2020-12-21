import React from 'react';

function Trio({column, children}) {
	return (
			<div className={`trio${column}`}>
				{children}
			</div>
	);
}

export default Trio;