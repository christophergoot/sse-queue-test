import React, { Component } from 'react';
import './App.css';
import ClaimSidebar from './ClaimSidebar';
import ActiveClaim from './ActiveClaim';
import axios from 'axios';

class App extends Component {
  constructor(props) {
    super(props);
    this.state = {
      activeClaim: false,
      allClaims: [{ id: '1', companyName: '', billedAmt: 0, active: true, changed: false, batchDate: '' }]
    }
  }
  eventSource = new EventSource('/api/stream');

  updateClaim = claimUpdates => {
    console.log(claimUpdates);
    return axios.post('/api/update-claim', claimUpdates)
  }
  handleClaimClose = claimId => {
    this.updateClaim({ id: claimId, active: false });
    this.setState({
      activeClaim: false,
    })
  }

  handleClaimControl = (claimId) => {
    const index = this.state.allClaims.findIndex(claim => claim.id === claimId);
    this.updateClaim({ id: claimId, active: true })
      .then(() => {
        this.setState({ 
          activeClaim: this.state.allClaims[index], 
        })    
      })
  }
  handleChanged = claimId => {
    const { activeClaim } = this.state;
    this.updateClaim({ id: claimId, changed: true })
      .then(() => {
        activeClaim.changed = true;
        this.setState({ activeClaim });
      })
      .catch(err => console.log('error hanging data', err))
  }
  updateAllClaims = (updatedClaim) => {
    const { allClaims } = this.state;
    const index = allClaims.findIndex(claim => claim.id === updatedClaim.id);
    allClaims[index] = { ...allClaims[index], ...updatedClaim };
    this.setState({ allClaims });
    console.log(`Claim from ${allClaims[index].companyName} is now ${allClaims[index].active ? 'Claimed Elsewhere' : 'Availiable to Claim'}`);
  }


  componentDidMount() {
    axios.get(`/api/claims`)
      .then(res => this.setState({ allClaims: res.data }));
    this.eventSource.addEventListener('claimUpdate', e => this.updateAllClaims(JSON.parse(e.data)) );
  }

  

  render() {
    return (
      <div className="App">
        <ClaimSidebar allClaims={this.state.allClaims} handleClaimControl={this.handleClaimControl} />
        <ActiveClaim handleClaimClose={this.handleClaimClose} activeClaim={this.state.activeClaim} handleChanged={this.handleChanged}/>
      </div>
    );
  }
}

export default App;
