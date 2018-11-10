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
    this.eventSource = new EventSource('http://localhost:8000/claims/');

    this.handleClaimClose = claimId => {
      const { allClaims } = this.state;
      const index = allClaims.findIndex(claim => claim.id === claimId);
      allClaims[index].active = false
      this.setState({
        activeClaim: false,
        allClaims
      })
    }
    this.handleClaimControl = (claimId) => {
      console.log(`claim ${claimId} is now active`);
      const index = this.state.allClaims.findIndex(claim => claim.id === claimId);
      const prevIndex = this.state.allClaims.findIndex(claim => claim.active === true);
      const allClaims = this.state.allClaims;
      allClaims[index].active = true;
      if (prevIndex >= 0) allClaims[prevIndex].active = false;
      this.setState({ 
        activeClaim: this.state.allClaims[index], 
        allClaims
      })
    }
    this.handleChanged = claimId => {
      const index = this.state.allClaims.findIndex(claim => claim.id === claimId);
      const { allClaims, activeClaim } = this.state;
      allClaims[index].changed = true;
      activeClaim.changed = true;
      this.setState({ allClaims, activeClaim })
    }
    this.updateAllClaims = (claimData) => {
      console.log('New data from eventSource', claimData);
    }
  }
  
  componentDidMount() {
    axios.get(`http://localhost:8000/claims/`)
      .then(res => res.data)
      .then(json => {
        this.setState({ allClaims: json })
      });
    this.eventSource.onmessage = e => this.updateAllClaims(JSON.parse(e.data));
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
