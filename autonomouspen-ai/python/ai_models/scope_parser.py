from langchain import OpenAI, LLMChain
from langchain.prompts import PromptTemplate
import json

class ScopeParser:
    def __init__(self):
        self.llm = OpenAI(temperature=0.0)
        self.scope_parser = self._create_scope_parser()

    def _create_scope_parser(self) -> LLMChain:
        prompt = PromptTemplate(
            input_variables=["policy_text"],
            template="""
            You are an expert at parsing bug bounty program policies.

            Given the following policy text:
            {policy_text}

            Extract the following information in JSON format:
            - in_scope: a list of in-scope domains, IPs, and wildcards
            - out_of_scope: a list of out-of-scope domains, IPs, and wildcards
            - special_rules: a list of any special rules, such as "no automated scanning" or "only these vulns"
            """
        )
        return LLMChain(llm=self.llm, prompt=prompt)

    def parse_scope(self, policy_text: str) -> dict:
        result = self.scope_parser.run(policy_text=policy_text)
        return json.loads(result)

if __name__ == "__main__":
    parser = ScopeParser()
    policy = """
    In-scope:
    - *.example.com
    - 192.168.1.0/24

    Out-of-scope:
    - *.staging.example.com
    - /admin/*

    Special Rules:
    - No automated scanning
    - Only SQLi and XSS vulnerabilities are in scope
    """
    scope = parser.parse_scope(policy)
    print(json.dumps(scope, indent=2))
