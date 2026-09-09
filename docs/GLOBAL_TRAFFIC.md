# FTN Global Traffic Integration

FTN cannot legitimately receive another platform's entire Internet traffic merely by naming that platform. Traffic reaches FTN only through a valid path: customer access, authorized transit, IX peering, an approved CDN/edge contract, or public telemetry.

## Target ecosystem

The integration catalog covers content and edge ecosystems such as Google, Facebook, Netflix, TikTok, IMO, PUBG, Free Fire, EdgeNext, Akamai, AWS, Cloudflare, Fastly, Bunny, and Tencent Cloud.

## Open-source building blocks

- Apache Traffic Server: scalable cache/proxy edge
- Varnish: HTTP caching and edge policy
- NGINX/OpenResty: reverse proxy and cache
- HAProxy/Envoy: proxy and load-balancing layers
- GoBGP/BIRD: authorized BGP control
- Prometheus/Grafana/OpenTelemetry: telemetry
- nfdump/pmacct/YA(F): flow collection and analysis

Apache Traffic Server and Apache Traffic Control are suitable references for a self-hosted CDN architecture. They do not create rights to third-party content or traffic.

## FTN acquisition modes

1. **Transit:** buy/establish authorized upstream connectivity.
2. **IX peering:** establish an approved peering session and exchange agreed routes.
3. **CDN:** serve FTN-owned or authorized content through edge caches.
4. **Partner edge:** integrate with a provider where the provider's contract/API permits it.
5. **Telemetry:** collect NetFlow/IPFIX/sFlow from FTN-controlled routers and approved sources.

No packet interception, credential reuse, route hijacking, or unauthorized access is part of this design.
