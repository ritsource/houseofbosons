import React, { useState, useEffect } from 'react';
import Rodal from 'rodal';
import { FaTrash, FaPen } from 'react-icons/fa';

import 'rodal/lib/rodal.css';

const TopicManageModal = (props) => {
	const [ valid, setValid ] = useState(false);
	const [ newTitle, setNewTitle ] = useState('');

	const [ loading, setLoading ] = useState(true);
	const [ errorMsg, setErrorMsg ] = useState('');

	const [ selected, setSelected ] = useState([]);

	const closeModal = () => {
		setValid(false);
		setLoading(false);
		setErrorMsg('');
		props.onClose();
	};

	useEffect(() => {
		(async function() {
			setSelected(props.alreadySelected.map((title) => ({ title })));

			try {
				await props.readTopics();
			} catch (error) {
				console.log(error);
				setErrorMsg(error.message);
			}
			setLoading(false);
		})();
	}, []);

	const onSelection = async () => {
		await setLoading(true);

		try {
			await props.onSelection(selected);
			closeModal();
		} catch (error) {
			setErrorMsg(error.message);
		}

		setLoading(false);
	};

	const deleteSelected = async () => {
		await setLoading(true);

		for (const t of selected) {
			try {
				await props.deleteTopic(t._id);
			} catch (error) {
				console.log(error);
				setErrorMsg(error.message);
				break;
			}
		}

		setLoading(false);
	};

	useEffect(
		() => {
			if (newTitle.length === 0) {
				setValid(false);
			}

			if (props.topics.length === 0 || newTitle.length > 0) {
				setValid(true);
				return;
			}
		},
		[ newTitle ]
	);

	return (
		<Rodal
			visible={props.visible}
			onClose={closeModal}
			showCloseButton={false}
			className="Rodal-ClassName"
			customStyles={{
				height: 'auto',
				width: 'auto',
				bottom: 'auto',
				top: 'auto',
				left: 'auto',
				right: 'auto',
				padding: '0px',
				backgroundColor: 'rgba(0,0,0,0)'
			}}
		>
			<div
				style={{
					width: '240px',
					background: 'white',
					borderRadius: '4px',
					padding: '20px'
				}}
			>
				<div
					style={{
						margin: '0px 0px 10px 0px',
						display: 'flex',
						justifyContent: 'space-between',
						alignItems: 'center'
					}}
				>
					<h4 style={{ margin: '0px', padding: '0px' }}>{props.text}</h4>
					{loading ? (
						<div className="Theme-Loading-Spin-Div" />
					) : props.forSelection ? (
						<button onClick={onSelection} className="Theme-Btn">
							Select
						</button>
					) : (
						<button onClick={deleteSelected} disabled={!selected.length} className="Theme-Btn">
							Delete
						</button>
					)}
				</div>
				<div
					id="Topics-Container-ID-0"
					style={{
						// button offset height 27px
						maxHeight: (27 + 10) * 6 + 'px',
						margin: '5px 0px 5px -5px',
						width: 'calc(100% + 5px)',
						overflowY: 'auto'
						// border: '1px solid red'
					}}
				>
					{props.topics.map((t) => {
						return (
							<TopicButton
								topic={t}
								loading={loading}
								deleteTopic={() => props.deleteTopic(t._id)}
								selected={selected.some(({ title }) => title === t.title)}
								addToSel={() => setSelected([ ...selected, t ])}
								removeFromSel={() => setSelected(selected.filter(({ title }) => title !== t.title))}
							/>
						);
					})}
				</div>
				<form
					onClick={async (e) => {
						e.preventDefault();

						if (!valid) {
							return;
						}

						const nt = newTitle.trim();
						if (props.topics.some(({ title }) => title === nt)) {
							setErrorMsg(`Title "${newTitle.trim()}" is already exist`);
							return;
						}

						setErrorMsg('');

						try {
							await props.createTopic(nt);
							setErrorMsg('');
							setNewTitle('');
							const el = document.querySelector('#Topics-Container-ID-0');
							el.scrollTop = el.scrollHeight;
						} catch (error) {
							console.log(error);
							setErrorMsg(error.message);
						}
					}}
				>
					<input
						placeholder="New topic title"
						style={{ width: 'calc(100% - 20px)' }}
						className="Theme-Input"
						value={newTitle}
						onChange={(e) => {
							setNewTitle(e.target.value);
						}}
					/>
					{!!errorMsg && (
						<p
							style={{
								color: '#ea4335',
								fontSize: '13px',
								margin: '10px 0px 0px 0px',
								padding: '0px'
							}}
						>
							{errorMsg}
						</p>
					)}
					<button style={{ marginTop: '10px' }} disabled={!valid} className="Theme-Btn" type="submit">
						Add
					</button>
				</form>
			</div>
		</Rodal>
	);
};

const TopicButton = (props) => {
	return (
		<button
			onClick={() => {
				if (props.loading) {
					return;
				}

				if (props.selected) {
					props.removeFromSel();
				} else {
					props.addToSel();
				}
			}}
			style={{
				textDecoration: props.selected ? 'underline' : 'none',
				margin: '5px',
				background: props.selected ? '#7EB3C7' : '#0090C3',
				border: '1px solid ' + props.selected ? '#7EB3C7' : '#0090C3'
			}}
			className="Theme-Btn"
		>
			{props.topic.title}
		</button>
	);
};

export default TopicManageModal;
