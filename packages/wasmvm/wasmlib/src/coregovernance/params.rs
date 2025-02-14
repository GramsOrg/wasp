// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
// >>>> DO NOT CHANGE THIS FILE! <<<<
// Change the schema definition file instead

#![allow(dead_code)]
#![allow(unused_imports)]

use crate::*;
use crate::coregovernance::*;

#[derive(Clone)]
pub struct ImmutableAddAllowedStateControllerAddressParams {
    pub(crate) proxy: Proxy,
}

impl ImmutableAddAllowedStateControllerAddressParams {
    pub fn new() -> ImmutableAddAllowedStateControllerAddressParams {
        ImmutableAddAllowedStateControllerAddressParams {
            proxy: params_proxy(),
        }
    }

    pub fn state_controller_address(&self) -> ScImmutableAddress {
        ScImmutableAddress::new(self.proxy.root(PARAM_STATE_CONTROLLER_ADDRESS))
    }
}

#[derive(Clone)]
pub struct MutableAddAllowedStateControllerAddressParams {
    pub(crate) proxy: Proxy,
}

impl MutableAddAllowedStateControllerAddressParams {
    pub fn state_controller_address(&self) -> ScMutableAddress {
        ScMutableAddress::new(self.proxy.root(PARAM_STATE_CONTROLLER_ADDRESS))
    }
}

#[derive(Clone)]
pub struct ImmutableAddCandidateNodeParams {
    pub(crate) proxy: Proxy,
}

impl ImmutableAddCandidateNodeParams {
    pub fn new() -> ImmutableAddCandidateNodeParams {
        ImmutableAddCandidateNodeParams {
            proxy: params_proxy(),
        }
    }

    pub fn access_node_info_access_api(&self) -> ScImmutableString {
        ScImmutableString::new(self.proxy.root(PARAM_ACCESS_NODE_INFO_ACCESS_API))
    }

    pub fn access_node_info_certificate(&self) -> ScImmutableBytes {
        ScImmutableBytes::new(self.proxy.root(PARAM_ACCESS_NODE_INFO_CERTIFICATE))
    }

    pub fn access_node_info_for_committee(&self) -> ScImmutableBool {
        ScImmutableBool::new(self.proxy.root(PARAM_ACCESS_NODE_INFO_FOR_COMMITTEE))
    }

    pub fn access_node_info_pub_key(&self) -> ScImmutableBytes {
        ScImmutableBytes::new(self.proxy.root(PARAM_ACCESS_NODE_INFO_PUB_KEY))
    }
}

#[derive(Clone)]
pub struct MutableAddCandidateNodeParams {
    pub(crate) proxy: Proxy,
}

impl MutableAddCandidateNodeParams {
    pub fn access_node_info_access_api(&self) -> ScMutableString {
        ScMutableString::new(self.proxy.root(PARAM_ACCESS_NODE_INFO_ACCESS_API))
    }

    pub fn access_node_info_certificate(&self) -> ScMutableBytes {
        ScMutableBytes::new(self.proxy.root(PARAM_ACCESS_NODE_INFO_CERTIFICATE))
    }

    pub fn access_node_info_for_committee(&self) -> ScMutableBool {
        ScMutableBool::new(self.proxy.root(PARAM_ACCESS_NODE_INFO_FOR_COMMITTEE))
    }

    pub fn access_node_info_pub_key(&self) -> ScMutableBytes {
        ScMutableBytes::new(self.proxy.root(PARAM_ACCESS_NODE_INFO_PUB_KEY))
    }
}

#[derive(Clone)]
pub struct MapBytesToImmutableUint8 {
    pub(crate) proxy: Proxy,
}

impl MapBytesToImmutableUint8 {
    pub fn get_uint8(&self, key: &[u8]) -> ScImmutableUint8 {
        ScImmutableUint8::new(self.proxy.key(&bytes_to_bytes(key)))
    }
}

#[derive(Clone)]
pub struct ImmutableChangeAccessNodesParams {
    pub(crate) proxy: Proxy,
}

impl ImmutableChangeAccessNodesParams {
    pub fn new() -> ImmutableChangeAccessNodesParams {
        ImmutableChangeAccessNodesParams {
            proxy: params_proxy(),
        }
    }

    pub fn change_access_nodes_actions(&self) -> MapBytesToImmutableUint8 {
        MapBytesToImmutableUint8 { proxy: self.proxy.root(PARAM_CHANGE_ACCESS_NODES_ACTIONS) }
    }
}

#[derive(Clone)]
pub struct MapBytesToMutableUint8 {
    pub(crate) proxy: Proxy,
}

impl MapBytesToMutableUint8 {
    pub fn clear(&self) {
        self.proxy.clear_map();
    }

    pub fn get_uint8(&self, key: &[u8]) -> ScMutableUint8 {
        ScMutableUint8::new(self.proxy.key(&bytes_to_bytes(key)))
    }
}

#[derive(Clone)]
pub struct MutableChangeAccessNodesParams {
    pub(crate) proxy: Proxy,
}

impl MutableChangeAccessNodesParams {
    pub fn change_access_nodes_actions(&self) -> MapBytesToMutableUint8 {
        MapBytesToMutableUint8 { proxy: self.proxy.root(PARAM_CHANGE_ACCESS_NODES_ACTIONS) }
    }
}

#[derive(Clone)]
pub struct ImmutableDelegateChainOwnershipParams {
    pub(crate) proxy: Proxy,
}

impl ImmutableDelegateChainOwnershipParams {
    pub fn new() -> ImmutableDelegateChainOwnershipParams {
        ImmutableDelegateChainOwnershipParams {
            proxy: params_proxy(),
        }
    }

    pub fn chain_owner(&self) -> ScImmutableAgentID {
        ScImmutableAgentID::new(self.proxy.root(PARAM_CHAIN_OWNER))
    }
}

#[derive(Clone)]
pub struct MutableDelegateChainOwnershipParams {
    pub(crate) proxy: Proxy,
}

impl MutableDelegateChainOwnershipParams {
    pub fn chain_owner(&self) -> ScMutableAgentID {
        ScMutableAgentID::new(self.proxy.root(PARAM_CHAIN_OWNER))
    }
}

#[derive(Clone)]
pub struct ImmutableRemoveAllowedStateControllerAddressParams {
    pub(crate) proxy: Proxy,
}

impl ImmutableRemoveAllowedStateControllerAddressParams {
    pub fn new() -> ImmutableRemoveAllowedStateControllerAddressParams {
        ImmutableRemoveAllowedStateControllerAddressParams {
            proxy: params_proxy(),
        }
    }

    pub fn state_controller_address(&self) -> ScImmutableAddress {
        ScImmutableAddress::new(self.proxy.root(PARAM_STATE_CONTROLLER_ADDRESS))
    }
}

#[derive(Clone)]
pub struct MutableRemoveAllowedStateControllerAddressParams {
    pub(crate) proxy: Proxy,
}

impl MutableRemoveAllowedStateControllerAddressParams {
    pub fn state_controller_address(&self) -> ScMutableAddress {
        ScMutableAddress::new(self.proxy.root(PARAM_STATE_CONTROLLER_ADDRESS))
    }
}

#[derive(Clone)]
pub struct ImmutableRevokeAccessNodeParams {
    pub(crate) proxy: Proxy,
}

impl ImmutableRevokeAccessNodeParams {
    pub fn new() -> ImmutableRevokeAccessNodeParams {
        ImmutableRevokeAccessNodeParams {
            proxy: params_proxy(),
        }
    }

    pub fn access_node_info_certificate(&self) -> ScImmutableBytes {
        ScImmutableBytes::new(self.proxy.root(PARAM_ACCESS_NODE_INFO_CERTIFICATE))
    }

    pub fn access_node_info_pub_key(&self) -> ScImmutableBytes {
        ScImmutableBytes::new(self.proxy.root(PARAM_ACCESS_NODE_INFO_PUB_KEY))
    }
}

#[derive(Clone)]
pub struct MutableRevokeAccessNodeParams {
    pub(crate) proxy: Proxy,
}

impl MutableRevokeAccessNodeParams {
    pub fn access_node_info_certificate(&self) -> ScMutableBytes {
        ScMutableBytes::new(self.proxy.root(PARAM_ACCESS_NODE_INFO_CERTIFICATE))
    }

    pub fn access_node_info_pub_key(&self) -> ScMutableBytes {
        ScMutableBytes::new(self.proxy.root(PARAM_ACCESS_NODE_INFO_PUB_KEY))
    }
}

#[derive(Clone)]
pub struct ImmutableRotateStateControllerParams {
    pub(crate) proxy: Proxy,
}

impl ImmutableRotateStateControllerParams {
    pub fn new() -> ImmutableRotateStateControllerParams {
        ImmutableRotateStateControllerParams {
            proxy: params_proxy(),
        }
    }

    pub fn state_controller_address(&self) -> ScImmutableAddress {
        ScImmutableAddress::new(self.proxy.root(PARAM_STATE_CONTROLLER_ADDRESS))
    }
}

#[derive(Clone)]
pub struct MutableRotateStateControllerParams {
    pub(crate) proxy: Proxy,
}

impl MutableRotateStateControllerParams {
    pub fn state_controller_address(&self) -> ScMutableAddress {
        ScMutableAddress::new(self.proxy.root(PARAM_STATE_CONTROLLER_ADDRESS))
    }
}

#[derive(Clone)]
pub struct ImmutableSetChainInfoParams {
    pub(crate) proxy: Proxy,
}

impl ImmutableSetChainInfoParams {
    pub fn new() -> ImmutableSetChainInfoParams {
        ImmutableSetChainInfoParams {
            proxy: params_proxy(),
        }
    }

    // default maximum size of a blob
    pub fn max_blob_size(&self) -> ScImmutableUint32 {
        ScImmutableUint32::new(self.proxy.root(PARAM_MAX_BLOB_SIZE))
    }

    // default maximum size of a single event
    pub fn max_event_size(&self) -> ScImmutableUint16 {
        ScImmutableUint16::new(self.proxy.root(PARAM_MAX_EVENT_SIZE))
    }

    // default maximum number of events per request
    pub fn max_events_per_req(&self) -> ScImmutableUint16 {
        ScImmutableUint16::new(self.proxy.root(PARAM_MAX_EVENTS_PER_REQ))
    }
}

#[derive(Clone)]
pub struct MutableSetChainInfoParams {
    pub(crate) proxy: Proxy,
}

impl MutableSetChainInfoParams {
    // default maximum size of a blob
    pub fn max_blob_size(&self) -> ScMutableUint32 {
        ScMutableUint32::new(self.proxy.root(PARAM_MAX_BLOB_SIZE))
    }

    // default maximum size of a single event
    pub fn max_event_size(&self) -> ScMutableUint16 {
        ScMutableUint16::new(self.proxy.root(PARAM_MAX_EVENT_SIZE))
    }

    // default maximum number of events per request
    pub fn max_events_per_req(&self) -> ScMutableUint16 {
        ScMutableUint16::new(self.proxy.root(PARAM_MAX_EVENTS_PER_REQ))
    }
}

#[derive(Clone)]
pub struct ImmutableSetFeePolicyParams {
    pub(crate) proxy: Proxy,
}

impl ImmutableSetFeePolicyParams {
    pub fn new() -> ImmutableSetFeePolicyParams {
        ImmutableSetFeePolicyParams {
            proxy: params_proxy(),
        }
    }

    pub fn fee_policy_bytes(&self) -> ScImmutableBytes {
        ScImmutableBytes::new(self.proxy.root(PARAM_FEE_POLICY_BYTES))
    }
}

#[derive(Clone)]
pub struct MutableSetFeePolicyParams {
    pub(crate) proxy: Proxy,
}

impl MutableSetFeePolicyParams {
    pub fn fee_policy_bytes(&self) -> ScMutableBytes {
        ScMutableBytes::new(self.proxy.root(PARAM_FEE_POLICY_BYTES))
    }
}
