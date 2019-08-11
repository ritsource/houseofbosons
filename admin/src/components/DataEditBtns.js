import React from 'react';

const DataEditBtns = (props) => {
	return (
		<React.Fragment>
			{props.editable ? (
				<div style={{ marginTop: '10px' }}>
					<button
						style={{ marginLeft: '0px' }}
						className="Theme-Btn Theme-Btn-Grey"
						onClick={() => props.setEditable(false)}
					>
						Cancel
					</button>
					<button style={{ marginLeft: '10px' }} className="Theme-Btn" onClick={props.onSave}>
						Save
					</button>
				</div>
			) : (
				<button style={{ marginTop: '10px' }} className="Theme-Btn" onClick={() => props.setEditable(true)}>
					Edit
				</button>
			)}
		</React.Fragment>
	);
};

export default DataEditBtns;
