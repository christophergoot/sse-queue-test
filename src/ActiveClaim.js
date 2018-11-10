import React, { Component } from 'react';

class ActiveClaim extends Component {
  render() {
    const { activeClaim: claim } = this.props;
    return !claim ? '' :
      <div key={claim.id} className='big claim' style={{ color: 'black', cursor: 'pointer', backgroundColor: claim.changed ? 'rgba(24, 191, 74, 0.19)' : 'inherit' }} >
        <p><span style={{ marginRight: '.5em', color: claim.changed ? 'green' : 'inherit' }}>{claim.changed ? '☑' : '☐'}</span>{claim.companyName}</p>
        <p>Billed Amount: ${claim.billedAmt}</p>
        <p>Date: {new Date(claim.batchDate).toDateString()}</p>
        <p>
          <input type='checkbox' id={claim.id} checked={claim.changed} onChange={() => this.props.handleChanged(claim.id)}/>
          <label htmlFor={claim.id}>Change the Data</label>
        </p>
        <button onClick={() => this.props.handleClaimClose(claim.id)} >Close Claim</button>
      </div>
  }
}

export default ActiveClaim;
