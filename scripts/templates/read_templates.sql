SELECT template_id, template_name, template_description, template_content 
FROM speech_templates
LIMIT $1 OFFSET $2;
