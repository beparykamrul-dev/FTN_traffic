# FTN Global Traffic Integration

FTN traffic acquisition is modeled as separate delivery, transit, and telemetry paths. A provider is never treated as a source of arbitrary third-party traffic merely because it is listed in the catalog.

## Target ecosystem

The integration catalog covers content and edge ecosystems such as Google, Facebook, Netflix, TikTok, IMO, PUBG, Free Fire, EdgeNext, Akamai, AWS, Cloudflare, Fastly, Bunny, and Tencent Cloud.

## Public distribution intermediaries

For FTN-owned or otherwise authorized public content, FTN can use public distribution services such as **jsDelivr, unpkg, cdnjs, GitHub Pages, and public package registries**. The pattern is:

`FTN Origin -> Public Distribution Intermediary -> End User`

These services can expand content delivery reach and reduce origin load. They are **not Internet-transit providers** and do not provide a mechanism for obtaining arbitrary private third-party traffic.

## Open-source building blocks

- Apache Traffic Server: scalable cache/proxy edge
- Varnish: HTTP caching and edge policy
- NGINX/OpenResty: reverse proxy and cache
- HAProxy/Envoy: proxy and load-balancing layers
- GoBGP/BIRD: authorized BGP control
- Prometheus/Grafana/OpenTelemetry: telemetry
- nfdump/pmacct/YAF: flow collection and analysis
- WebTorrent: distributed delivery for content FTN owns or is authorized to distribute

Apache Traffic Server and Apache Traffic Control are suitable references for a self-hosted CDN architecture. They do not create rights to third-party content or traffic.

## FTN acquisition modes

1. **Direct origin:** FTN origin to FTN edge/users.
2. **Public CDN/intermediary:** public or authorized FTN content through services such as jsDelivr-like distribution layers.
3. **Open/P2P distribution:** authorized content through distributed delivery.
4. **Transit:** buy/establish authorized upstream Internet connectivity.
5. **IX peering:** establish approved peering and exchange agreed routes.
6. **CDN contract:** integrate an authorized CDN/edge service.
7. **Partner edge:** use provider APIs/contracts where permitted.
8. **Telemetry:** collect NetFlow/IPFIX/sFlow from FTN-controlled routers and approved sources.

Route mutations, provider credentials, and privileged edge changes remain approval-gated. No packet interception, credential reuse, route hijacking, or unauthorized access is part of this design.
