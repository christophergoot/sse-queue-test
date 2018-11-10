import React, { Component } from 'react';

class ClaimSidebar extends Component {
  renderClaims(claimData) {
    return claimData.map(claim =>
      <div key={claim.id} className='small claim' style={claim.active ? { color: 'lightgrey', cursor: 'not-allowed' } : { color: 'black', cursor: 'pointer', backgroundColor: claim.changed ? 'rgba(24, 191, 74, 0.19)' : 'inherit' }} onClick={!claim.active ? () => this.props.handleClaimControl(claim.id) : null} >
        <span style={{ marginRight: '.5em', color: claim.changed ? 'green' : 'inherit' }}>{claim.changed ? '☑' : '☐'}</span>{claim.companyName}: ${claim.billedAmt}
      </div>
    )
  }

  render() {
    const { allClaims } = this.props;
    return (
      <div>
        {this.renderClaims(allClaims)}
      </div>
    );
  }
}


export default ClaimSidebar;