"""
AUTONOMOUSPEN AI - Decision Engine
Fully autonomous penetration testing brain
"""

from typing import List, Dict, Any
import json
from langchain import OpenAI, LLMChain
from langchain.prompts import PromptTemplate

class AutonomousDecisionEngine:
  """Core AI brain for autonomous penetration testing"""

  def __init__(self):
      self.llm = OpenAI(temperature=0.7)
      self.scan_planner = self._create_scan_planner()
      self.payload_generator = self._create_payload_generator()

  def _create_scan_planner(self) -> LLMChain:
      """Creates the AI chain for planning scans"""
      prompt = PromptTemplate(
          input_variables=["recon_data", "technology_stack"],
          template="""
You are AUTONOMOUSPEN AI, an expert autonomous penetration tester.

Reconnaissance Data:
{recon_data}

Technology Stack:
{technology_stack}

Create a prioritized testing plan with:

* vulnerability_priorities (ordered by impact)
* test_sequences (specific attack chains)
* time_allocation (minutes per test)
* bypass_strategies (for detected protections)
Output as JSON.
"""
      )
      return LLMChain(llm=self.llm, prompt=prompt)

  def _create_payload_generator(self) -> LLMChain:
      """Creates the AI chain for generating payloads"""
      prompt = PromptTemplate(
          input_variables=["vulnerability_type", "context", "waf_detected"],
          template="""
Generate advanced payloads for {vulnerability_type}.

Context: {context}
WAF Detected: {waf_detected}

Create 10 payloads with increasing sophistication and evasion techniques.
Include polyglots and encoding variations.

Output as JSON array.
"""
      )
      return LLMChain(llm=self.llm, prompt=prompt)

  def plan_scan(self, target: Dict[str, Any]) -> Dict[str, Any]:
      """Plans the scanning strategy for a target"""
      result = self.scan_planner.run(
          recon_data=json.dumps(target.get("recon", {})),
          technology_stack=target.get("tech_stack", "Unknown")
      )
      return json.loads(result)

  def generate_payloads(self, vuln_type: str, context: str, waf: bool = False) -> List[str]:
      """Generates context-aware payloads"""
      result = self.payload_generator.run(
          vulnerability_type=vuln_type,
          context=context,
          waf_detected=str(waf)
      )
      return json.loads(result)

if __name__ == "__main__":
  # Test the engine
  engine = AutonomousDecisionEngine()
  print("✅ AUTONOMOUSPEN AI Decision Engine initialized")
