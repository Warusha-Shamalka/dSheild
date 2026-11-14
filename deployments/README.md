# Deployment configuration

This folder will contain configuration for packaging and running dSheild itself.
It is separate from the example infrastructure files that the scanner reads in
[tests](../tests/).

Planned contents, introduced as the corresponding roadmap phases are reached:

- **Phase 6:** Docker build configuration and Docker Compose definitions for the
  application, PostgreSQL, and Redis. A build file can live at the repository root
  if that makes its build context clearer; document the choice here.
- **Phase 10:** Kubernetes manifests for Deployments, Services, configuration,
  probes, and service accounts.
- **Phase 11:** admission-webhook deployment and certificate configuration.
- **Phase 12:** Argo CD application definitions for GitOps.
- **Phase 13:** Prometheus scrape configuration and Grafana dashboard definitions.
- **Phase 14:** Terraform for dSheild's own infrastructure.

**Current phase:** this folder is a placeholder. The Dockerfile parser reads text;
it does not need Docker, a running cluster, or deployed cloud infrastructure.
Store only example configuration here; inject real credentials at runtime.
