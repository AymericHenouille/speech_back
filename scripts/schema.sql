DROP TABLE IF EXISTS speech_templates;
DROP SEQUENCE IF EXISTS speech_template_id_seq;

CREATE SEQUENCE IF NOT EXISTS speech_template_id_seq;

CREATE TABLE IF NOT EXISTS speech_templates (
  template_id INT DEFAULT nextval('speech_template_id_seq'),
  template_name VARCHAR(255) NOT NULL,
  template_description TEXT NOT NULL,
  template_content TEXT NOT NULL,
  template_insert_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  CONSTRAINT speech_templates_pk PRIMARY KEY (template_id)
);

INSERT INTO speech_templates (template_name, template_description, template_content)
SELECT
  'Template ' || i,
  'Description for template ' || i,
  'Content for template ' || i
FROM generate_series(1, 1000) i;